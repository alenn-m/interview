package service

import (
	"context"

	"github.com/alenn-m/interview/svc/pkg/pack/entity"
	"github.com/alenn-m/interview/svc/pkg/pack/repository"
)

type Service interface {
	Create(ctx context.Context, pack *entity.Pack) (*entity.Pack, error)
	Update(ctx context.Context, pack *entity.Pack) (*entity.Pack, error)
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (*entity.Pack, error)
	List(ctx context.Context) ([]*entity.Pack, error)
}

type service struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, pack *entity.Pack) (*entity.Pack, error) {
	return s.repo.Create(ctx, pack)
}

func (s *service) Update(ctx context.Context, pack *entity.Pack) (*entity.Pack, error) {
	return s.repo.Update(ctx, pack)
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *service) Get(ctx context.Context, id int) (*entity.Pack, error) {
	return s.repo.Get(ctx, id)
}

func (s *service) List(ctx context.Context) ([]*entity.Pack, error) {
	return s.repo.List(ctx)
}
