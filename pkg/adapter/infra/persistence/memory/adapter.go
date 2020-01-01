package memory

import (
	"github.com/minchao/go-realworld/pkg/application/article/port"
)

type Adapter struct {
	articleRepository *ArticleRepository
	tagRepository     *TagRepository
}

func NewAdapter() (*Adapter, error) {
	return &Adapter{
		articleRepository: NewArticleRepository(),
		tagRepository:     NewTagRepository(),
	}, nil
}

func (a *Adapter) Article() port.ArticleRepository {
	return a.articleRepository
}

func (a *Adapter) Tag() port.TagRepository {
	return a.tagRepository
}
