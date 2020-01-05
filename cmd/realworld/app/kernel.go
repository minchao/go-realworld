package app

import (
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/minchao/go-realworld/pkg/adapter/infra/persistence/memory"
	"github.com/minchao/go-realworld/pkg/adapter/ui/rest"
	"github.com/minchao/go-realworld/pkg/application/article"
	articlePort "github.com/minchao/go-realworld/pkg/application/article/port"
)

type Kernel struct {
	Logger log.Logger
	Router *mux.Router

	articleRepository articlePort.ArticleRepository
	tagRepository     articlePort.TagRepository

	articleService article.UseCase
	tagService     article.TagUseCase
}

func NewKernel() *Kernel {
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	return &Kernel{
		Logger: logger,
		Router: mux.NewRouter(),
	}
}

func (k *Kernel) Boot() error {
	for _, fn := range []func() error{
		k.initializePersistence,
		k.initializeServices,
		k.initializeREST,
	} {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}

func (k *Kernel) initializePersistence() error {
	adapter, err := memory.NewAdapter()
	if err != nil {
		return err
	}

	k.articleRepository = adapter.Article()
	k.tagRepository = adapter.Tag()
	return nil
}

func (k *Kernel) initializeServices() error {
	k.tagService = article.NewTagService(k.tagRepository)
	k.articleService = article.NewArticleService(k.articleRepository, k.tagService)
	return nil
}

func (k *Kernel) initializeREST() error {
	httpLogger := log.With(k.Logger, "component", "HTTP")
	options := []httpTransport.ServerOption{
		httpTransport.ServerErrorHandler(transport.NewLogErrorHandler(httpLogger)),
		httpTransport.ServerErrorEncoder(rest.EncodeError),
	}

	rest.InitArticleHandler(k.Router, k.articleService, options)
	rest.InitTagHandler(k.Router, k.tagService, options)
	return nil
}
