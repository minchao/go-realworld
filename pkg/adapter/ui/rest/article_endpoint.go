package rest

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"

	"github.com/minchao/go-realworld/pkg/application/article"
	"github.com/minchao/go-realworld/pkg/application/article/domain"
	userDomain "github.com/minchao/go-realworld/pkg/application/user/domain"
)

type ArticleEndpoints struct {
	GetArticles   endpoint.Endpoint
	PostArticles  endpoint.Endpoint
	GetArticle    endpoint.Endpoint
	PutArticle    endpoint.Endpoint
	DeleteArticle endpoint.Endpoint
}

func makeArticleServerEndpoints(s article.UseCase) ArticleEndpoints {
	return ArticleEndpoints{
		GetArticles:   makeGetArticlesEndpoint(s),
		PostArticles:  makePostArticleEndpoint(s),
		GetArticle:    makeGetArticleEndpoint(s),
		PutArticle:    makePutArticleEndpoint(s),
		DeleteArticle: nil,
	}
}

type articleData struct {
	Slug           string          `json:"slug"`
	Title          string          `json:"title"`
	Description    string          `json:"description"`
	Body           string          `json:"body"`
	TagList        []domain.Tag    `json:"tagList"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      time.Time       `json:"updatedAt"`
	Favorited      bool            `json:"favorited"`
	FavoritesCount int             `json:"favoritesCount"`
	Author         userDomain.User `json:"author"`
}

func transformArticle(article domain.Article) articleData {
	return articleData{
		Slug:           article.Slug,
		Title:          article.Title,
		Description:    article.Description,
		Body:           article.Body,
		TagList:        article.TagList,
		CreatedAt:      article.CreatedAt,
		UpdatedAt:      article.UpdatedAt,
		Favorited:      article.Favorited,
		FavoritesCount: article.FavoritesCount,
		Author:         article.Author,
	}
}

func transformArticles(articles []*domain.Article) []articleData {
	as := make([]articleData, len(articles))
	for i, v := range articles {
		as[i] = transformArticle(*v)
	}
	return as
}

type getArticlesResponse struct {
	Articles      []articleData `json:"articles"`
	ArticlesCount int           `json:"articlesCount"`
}

func makeGetArticlesEndpoint(s article.UseCase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		articles, err := s.GetAll(ctx, domain.ArticleCriteria{})
		if err != nil {
			return nil, err
		}
		return getArticlesResponse{
			Articles:      transformArticles(articles),
			ArticlesCount: len(articles),
		}, nil
	}
}

type postArticle struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Body        string   `json:"body"`
	TagList     []string `json:"tagList"`
}

type postArticleRequest struct {
	Article postArticle `json:"article"`
}

func (r postArticleRequest) toArticleCreationData(user userDomain.User) article.CreationData {
	return article.CreationData{
		Title:       r.Article.Title,
		Description: r.Article.Description,
		Body:        r.Article.Body,
		TagList:     r.Article.TagList,
		Author:      user,
	}
}

type postArticleResponse struct {
	Article articleData `json:"article"`
}

func makePostArticleEndpoint(service article.UseCase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postArticleRequest)
		newArticle, err := service.Create(ctx, req.toArticleCreationData(userDomain.User{}))
		if err != nil {
			return nil, err
		}
		return postArticleResponse{Article: transformArticle(*newArticle)}, nil
	}
}

type getArticleRequest struct {
	Slug string
}

type getArticleResponse struct {
	Article articleData `json:"article"`
}

func makeGetArticleEndpoint(s article.UseCase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getArticleRequest)
		a, err := s.Get(ctx, req.Slug)
		if err != nil {
			return nil, err
		}
		return getArticleResponse{Article: transformArticle(*a)}, nil
	}
}

type putArticle struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Body        *string `json:"body"`
}

type putArticleRequest struct {
	Slug    string     `json:"-"`
	Article putArticle `json:"article"`
}

func (r putArticleRequest) toArticleMutationData() article.MutationData {
	return article.MutationData{
		Title:       r.Article.Title,
		Description: r.Article.Description,
		Body:        r.Article.Body,
	}
}

type putArticleResponse struct {
	Article articleData `json:"article"`
}

func makePutArticleEndpoint(s article.UseCase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(putArticleRequest)
		newArticle, err := s.Update(ctx, req.Slug, req.toArticleMutationData())
		if err != nil {
			return nil, err
		}
		return putArticleResponse{Article: transformArticle(*newArticle)}, nil
	}
}
