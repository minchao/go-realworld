package cmd

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/minchao/go-realworld/cmd/realworld/app"
)

const (
	defaultServePort = 8080
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Starts the server",
		RunE:  serveRun,
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().Int("port", defaultServePort, "serve port")
}

func serveRun(_ *cobra.Command, _ []string) error {
	kernel := app.NewKernel()
	if err := kernel.Boot(); err != nil {
		return err
	}

	// TODO graceful shutdown
	httpAddr := fmt.Sprintf(":%d", defaultServePort)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		_ = kernel.Logger.Log("transport", "HTTP", "addr", httpAddr)
		errs <- http.ListenAndServe(httpAddr, kernel.Router)
	}()

	_ = kernel.Logger.Log("exit", <-errs)

	return nil
}
