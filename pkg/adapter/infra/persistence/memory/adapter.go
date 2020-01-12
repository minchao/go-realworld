package memory

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

func (a *Adapter) Article() *ArticleRepository {
	return a.articleRepository
}

func (a *Adapter) Tag() *TagRepository {
	return a.tagRepository
}
