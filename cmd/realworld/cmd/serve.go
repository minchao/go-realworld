package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"

	"github.com/minchao/go-realworld/cmd/realworld/app"
)

const (
	defaultServePort = 8080
)

var (
	port int

	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Starts the server",
		RunE:  serveRun,
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().IntVarP(&port, "port", "p", defaultServePort, "serve port")
}

func serveRun(_ *cobra.Command, _ []string) error {
	kernel := app.NewKernel()
	if err := kernel.Boot(); err != nil {
		return err
	}

	httpAddr := fmt.Sprintf(":%d", port)

	// see https://github.com/gorilla/mux#graceful-shutdown
	server := &http.Server{
		Handler:      kernel.Router,
		Addr:         httpAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		_ = kernel.Logger.Log("transport", "HTTP", "addr", httpAddr)

		if err := server.ListenAndServe(); err != nil {
			_ = kernel.Logger.Log("terminated", err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_ = server.Shutdown(ctx)
	_ = kernel.Logger.Log("terminated", "shutting down")
	return nil
}
