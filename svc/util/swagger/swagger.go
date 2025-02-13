package swagger

import (
	"net/http"

	"github.com/alenn-m/interview/svc/util/router"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Register adds swagger documentation to the router
func Register(r router.Router) {
	r.Chi.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/api/swagger.yaml"),
	))

	// Serve the swagger.yaml file
	r.Chi.Get("/api/swagger.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "api/swagger.yaml")
	})
}
