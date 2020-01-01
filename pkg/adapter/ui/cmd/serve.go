package cmd

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	"github.com/minchao/go-realworld/pkg/adapter/infra/persistence/memory"
	"github.com/minchao/go-realworld/pkg/adapter/ui/rest"
	"github.com/minchao/go-realworld/pkg/application/article"
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
}

func serveRun(_ *cobra.Command, _ []string) error {
	httpAddr := ":8080"

	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	repositoryManager, _ := memory.NewAdapter()

	router := mux.NewRouter()
	httpLogger := log.With(logger, "component", "HTTP")
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(httpLogger)),
		httptransport.ServerErrorEncoder(rest.EncodeError),
	}

	tagService := article.NewTagService(repositoryManager.Tag())

	rest.InitArticleHandler(router, article.NewArticleService(repositoryManager.Article(), tagService), options)
	rest.InitTagHandler(router, tagService, options)

	// TODO graceful shutdown

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		_ = logger.Log("transport", "HTTP", "addr", httpAddr)
		errs <- http.ListenAndServe(httpAddr, router)
	}()

	_ = logger.Log("exit", <-errs)
	return nil
}
