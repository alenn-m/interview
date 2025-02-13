package pack

import (
	"github.com/alenn-m/interview/svc/pkg/pack/repository"
	"github.com/alenn-m/interview/svc/pkg/pack/service"
	"github.com/alenn-m/interview/svc/pkg/pack/transport/http"
	"go.uber.org/fx"
)

// Module exports the pack module
var Module = fx.Options(
	fx.Provide(
		repository.New,
		service.New,
		http.NewDecodeEncoder,
		NewClient,
	),
	fx.Invoke(
		http.Register,
	),
)
