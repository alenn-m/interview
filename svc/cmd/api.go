package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/alenn-m/interview/svc/pkg/order"
	"github.com/alenn-m/interview/svc/pkg/pack"
	"github.com/alenn-m/interview/svc/util/db"
	"github.com/alenn-m/interview/svc/util/logs"
	"github.com/alenn-m/interview/svc/util/router"
	"github.com/alenn-m/interview/svc/util/swagger"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		router.NewRouter,
		logs.GetLogger,
		db.NewDatabase,
	),
	order.Module,
	pack.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	logger logs.Logger,
	r router.Router,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Starting Application on port: ", os.Getenv("PORT"))

			// Register swagger documentation
			swagger.Register(r)

			go func() {
				http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r.Chi)
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Info("Stopping Application")
			return nil
		},
	})
}

var apiCommand = &cobra.Command{
	Use:   "api",
	Short: "Serve API",
	RunE: func(cmd *cobra.Command, args []string) error {
		os.Setenv("CONFIG_LOCATION", cfgFile)
		fx.New(Module).Run()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(apiCommand)
}
