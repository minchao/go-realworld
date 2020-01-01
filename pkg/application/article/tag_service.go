package article

import (
	"context"

	"github.com/minchao/go-realworld/pkg/application/article/domain"
	"github.com/minchao/go-realworld/pkg/application/article/port"
)

type TagService struct {
	repository port.TagRepository
}

func NewTagService(repo port.TagRepository) *TagService {
	return &TagService{
		repository: repo,
	}
}

func (s *TagService) Create(ctx context.Context, tag domain.Tag) error {
	return s.repository.Create(ctx, tag)
}

func (s *TagService) FindAll(ctx context.Context) ([]domain.Tag, error) {
	return s.repository.FindAll(ctx)
}
