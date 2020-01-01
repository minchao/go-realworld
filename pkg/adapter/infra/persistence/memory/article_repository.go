package memory

import (
	"context"
	"errors"
	"sync"

	"github.com/minchao/go-realworld/pkg/application/article/domain"
)

type ArticleRepository struct {
	sync.RWMutex
	articles map[string]*domain.Article
	list     []*domain.Article
}

func NewArticleRepository() *ArticleRepository {
	return &ArticleRepository{
		articles: map[string]*domain.Article{},
		list:     []*domain.Article{},
	}
}

func (r *ArticleRepository) Find(_ context.Context, slug string) (*domain.Article, error) {
	r.RLock()
	defer r.RUnlock()
	article, found := r.articles[slug]
	if !found {
		return nil, domain.ErrArticleNotFound
	}
	return article, nil
}

func (r *ArticleRepository) FindBy(_ context.Context, criteria domain.ArticleCriteria) ([]*domain.Article, error) {
	r.RLock()
	defer r.RUnlock()
	return r.list, nil
}

func (r *ArticleRepository) Save(_ context.Context, article *domain.Article) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.articles[article.Slug]; !ok {
		r.list = append(r.list, article)
	}
	r.articles[article.Slug] = article
	return nil
}

func (r *ArticleRepository) Delete(_ context.Context, slug string) error {
	r.Lock()
	defer r.Unlock()
	delete(r.articles, slug)
	if idx, err := r.findIndex(slug); err == nil {
		r.removeFromList(idx)
	}
	return nil
}

func (r *ArticleRepository) findIndex(slug string) (int, error) {
	for i, article := range r.list {
		if slug == article.Slug {
			return i, nil
		}
	}
	return 0, errors.New("not found")
}

// see https://yourbasic.org/golang/delete-element-slice/
func (r *ArticleRepository) removeFromList(idx int) {
	copy(r.list[idx:], r.list[idx+1:])
	r.list[len(r.list)-1] = nil
	r.list = r.list[:len(r.list)-1]
}
