package port

import (
	"context"

	"github.com/minchao/go-realworld/pkg/application/article/domain"
)

type ArticleRepository interface {
	Find(ctx context.Context, slug string) (*domain.Article, error)
	FindBy(ctx context.Context, criteria domain.ArticleCriteria) ([]*domain.Article, error)
	Save(ctx context.Context, article *domain.Article) error
	Delete(ctx context.Context, slug string) error
}

type CreateTagService interface {
	Create(ctx context.Context, tag domain.Tag) error
}
