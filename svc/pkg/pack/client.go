package pack

import (
	"context"

	"github.com/alenn-m/interview/svc/pkg/pack/entity"
	"github.com/alenn-m/interview/svc/pkg/pack/service"
	"go.uber.org/fx"
)

type Client interface {
	List(ctx context.Context) ([]*entity.Pack, error)
}

type ClientOptions struct {
	fx.In

	Service service.Service
}

type client struct {
	service service.Service
}

func (c client) List(ctx context.Context) ([]*entity.Pack, error) {
	return c.service.List(ctx)
}

func NewClient(opts ClientOptions) Client {
	return &client{service: opts.Service}
}
