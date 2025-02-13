package order

import (
	"github.com/alenn-m/interview/svc/pkg/order/service"
	"github.com/alenn-m/interview/svc/pkg/order/transport/http"
	"go.uber.org/fx"
)

// Module exports the order module
var Module = fx.Options(
	fx.Provide(
		service.New,
		http.NewDecodeEncoder,
	),
	fx.Invoke(
		http.Register,
	),
)
