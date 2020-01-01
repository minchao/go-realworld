package memory

import (
	"context"
	"sync"

	"github.com/minchao/go-realworld/pkg/application/article/domain"
)

type TagRepository struct {
	sync.RWMutex
	tags []domain.Tag
}

func NewTagRepository() *TagRepository {
	return &TagRepository{
		tags: []domain.Tag{},
	}
}

func (r *TagRepository) Create(_ context.Context, tag domain.Tag) error {
	r.Lock()
	defer r.Unlock()
	for _, v := range r.tags {
		if tag == v {
			return nil
		}
	}
	r.tags = append(r.tags, tag)
	return nil
}

func (r *TagRepository) FindAll(_ context.Context) ([]domain.Tag, error) {
	r.RLock()
	tags := append(r.tags[:0:0], r.tags...)
	defer r.RUnlock()
	return tags, nil
}
