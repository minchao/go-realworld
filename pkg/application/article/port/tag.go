package port

import (
	"context"

	"github.com/minchao/go-realworld/pkg/application/article/domain"
)

type TagRepository interface {
	Create(ctx context.Context, tag domain.Tag) error
	FindAll(ctx context.Context) ([]domain.Tag, error)
}
