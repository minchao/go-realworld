package article

import (
	"context"

	"github.com/minchao/go-realworld/pkg/application/article/domain"
)

type TagUseCase interface {
	FindAll(ctx context.Context) ([]domain.Tag, error)
}
