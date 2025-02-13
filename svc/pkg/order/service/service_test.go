package service

import (
	"context"
	"testing"

	"github.com/alenn-m/interview/svc/pkg/order/entity"
	packEntity "github.com/alenn-m/interview/svc/pkg/pack/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPackClient is a mock implementation of PackClient interface
type MockPackClient struct {
	mock.Mock
}

func (m *MockPackClient) List(ctx context.Context) ([]*packEntity.Pack, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*packEntity.Pack), args.Error(1)
}

func TestService_Create(t *testing.T) {
	tests := []struct {
		name     string
		request  *entity.Request
		expected map[int]int // map[packSize]count
	}{
		{
			name: "1 item",
			request: &entity.Request{
				ItemsNumber: 1,
			},
			expected: map[int]int{
				250: 1,
			},
		},
		{
			name: "250 items",
			request: &entity.Request{
				ItemsNumber: 250,
			},
			expected: map[int]int{
				250: 1,
			},
		},
		{
			name: "251 items",
			request: &entity.Request{
				ItemsNumber: 251,
			},
			expected: map[int]int{
				500: 1,
			},
		},
		{
			name: "501 items",
			request: &entity.Request{
				ItemsNumber: 501,
			},
			expected: map[int]int{
				500: 1,
				250: 1,
			},
		},
		{
			name: "12001 items",
			request: &entity.Request{
				ItemsNumber: 12001,
			},
			expected: map[int]int{
				5000: 2,
				2000: 1,
				250:  1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock pack client
			mockPackClient := new(MockPackClient)
			// Setup expectation - return the standard pack sizes
			mockPackClient.On("List", mock.Anything).Return([]*packEntity.Pack{
				{Amount: 5000},
				{Amount: 2000},
				{Amount: 1000},
				{Amount: 500},
				{Amount: 250},
			}, nil)

			// Create service with mock pack client
			svc := New(Options{
				PackClient: mockPackClient,
			})

			result, err := svc.Create(context.Background(), tt.request)
			assert.NoError(t, err)

			// Verify pack client was called
			mockPackClient.AssertExpectations(t)

			// Convert result to map[packSize]count for easier comparison
			resultMap := make(map[int]int)
			for _, pack := range result.Packs {
				resultMap[pack.Amount] = pack.Count
			}

			assert.Equal(t, tt.expected, resultMap)

			// Verify total items is sufficient but not excessive
			if tt.request.ItemsNumber > 0 {
				// Check that we're providing enough items
				assert.GreaterOrEqual(t, result.TotalItems, tt.request.ItemsNumber)

				// Check that we're not using more packs than necessary
				// by verifying that removing any pack would make the total insufficient
				for _, pack := range result.Packs {
					remainingItems := result.TotalItems - pack.Amount
					if pack.Count > 1 {
						remainingItems = result.TotalItems - pack.Amount*pack.Count
					}
					assert.Less(t, remainingItems, tt.request.ItemsNumber,
						"Solution uses more packs than necessary")
				}
			}
		})
	}
}
