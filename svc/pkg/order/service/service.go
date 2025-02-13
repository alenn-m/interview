package service

import (
	"context"
	"fmt"
	"math"
	"sort"

	"github.com/alenn-m/interview/svc/pkg/order/entity"
	"github.com/alenn-m/interview/svc/pkg/pack"
	"go.uber.org/fx"
)

type Service interface {
	Create(ctx context.Context, req *entity.Request) (*entity.Response, error)
}

type Options struct {
	fx.In

	PackClient pack.Client
}

type service struct {
	packClient pack.Client
}

func New(o Options) Service {
	return &service{packClient: o.PackClient}
}

func (s *service) Create(ctx context.Context, req *entity.Request) (*entity.Response, error) {
	if req.ItemsNumber <= 0 {
		return &entity.Response{
			ItemsNumber: req.ItemsNumber,
			TotalItems:  0,
			Packs:       []entity.PackResponse{},
		}, nil
	}

	packItems, err := s.packClient.List(ctx)
	if err != nil {
		return nil, err
	}

	var packSizes []int
	for _, p := range packItems {
		packSizes = append(packSizes, p.Amount)
	}

	// Find the optimal combination of packs
	packs := findOptimalPacks(packSizes, req.ItemsNumber)

	// Convert the result to response format
	response := &entity.Response{
		ItemsNumber: req.ItemsNumber,
		Packs:       make([]entity.PackResponse, 0, len(packs)),
	}

	// Create response with pack sizes and their counts
	for amount, count := range packs {
		if count > 0 {
			response.Packs = append(response.Packs, entity.PackResponse{
				Name:   fmt.Sprintf("%d pack", amount),
				Amount: amount,
				Count:  count,
			})
		}
	}

	// Sort packs by size in descending order for consistent response
	sort.Slice(response.Packs, func(i, j int) bool {
		return response.Packs[i].Amount > response.Packs[j].Amount
	})

	// Calculate total items
	response.CalculateTotalItems()

	return response, nil
}

// findOptimalPacks finds the optimal combination of packs for the given number of items
func findOptimalPacks(packSizes []int, items int) map[int]int {
	if items <= 0 {
		return map[int]int{}
	}

	// First try: find exact match
	for i := len(packSizes) - 1; i >= 0; i-- {
		if packSizes[i] == items {
			return map[int]int{packSizes[i]: 1}
		}
	}

	// Try all possible combinations starting from each pack size
	var bestSolution map[int]int
	minExcess := -1         // Track minimum excess items
	minPacks := math.MaxInt // Track minimum number of packs

	// First, try single packs that could work
	for _, size := range packSizes {
		if size > items {
			excess := size - items
			if minExcess == -1 || excess < minExcess {
				bestSolution = map[int]int{size: 1}
				minExcess = excess
				minPacks = 1
			}
		}
	}

	// Try combinations of packs
	for startIdx := 0; startIdx < len(packSizes); startIdx++ {
		solution := make(map[int]int)
		remainingItems := items
		totalPacks := 0

		// Try filling from this pack size down
		for i := startIdx; i < len(packSizes); i++ {
			size := packSizes[i]
			if remainingItems >= size {
				count := remainingItems / size
				solution[size] = count
				totalPacks += count
				remainingItems -= count * size
			}
		}

		// If we have remaining items, try to find the smallest pack that can handle them
		if remainingItems > 0 {
			for i := len(packSizes) - 1; i >= 0; i-- {
				if packSizes[i] >= remainingItems {
					solution[packSizes[i]]++
					totalPacks++
					remainingItems = 0
					break
				}
			}
		}

		if remainingItems > 0 {
			continue // Skip invalid solutions
		}

		// Calculate total items and excess
		totalItems := 0
		for size, count := range solution {
			totalItems += size * count
		}
		excess := totalItems - items

		// Update the best solution if:
		// 1. This is our first valid solution (minExcess < 0)
		// 2. This solution has less excess
		// 3. This solution has same excess but fewer packs
		if minExcess < 0 || excess < minExcess || (excess == minExcess && totalPacks < minPacks) {
			bestSolution = make(map[int]int)
			for k, v := range solution {
				bestSolution[k] = v
			}
			minExcess = excess
			minPacks = totalPacks
		}
	}

	return bestSolution
}
