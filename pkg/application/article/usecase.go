package article

import (
	"context"

	"github.com/minchao/go-realworld/pkg/application/article/domain"
)

type UseCase interface {
	Create(ctx context.Context, options CreateArticleOptions) (*domain.Article, error)
	Get(ctx context.Context, slug string) (*domain.Article, error)
	GetAll(ctx context.Context, criteria domain.ArticleCriteria) ([]*domain.Article, error)
	Update(ctx context.Context, article *domain.Article) error
	Delete(ctx context.Context, slug string) error
}
