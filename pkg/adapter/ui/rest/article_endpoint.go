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
	GetArticle    endpoint.Endpoint
	PostArticle   endpoint.Endpoint
	PutArticle    endpoint.Endpoint
	DeleteArticle endpoint.Endpoint
}

func makeArticleServerEndpoints(s article.UseCase) ArticleEndpoints {
	return ArticleEndpoints{
		GetArticles:   makeGetArticlesEndpoint(s),
		GetArticle:    makeGetArticleEndpoint(s),
		PostArticle:   makePostArticleEndpoint(s),
		PutArticle:    nil,
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

type articleOptions struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Body        string   `json:"body"`
	TagList     []string `json:"tagList"`
}

type postArticleRequest struct {
	Article articleOptions `json:"article"`
}

type postArticleResponse struct {
	Article articleData `json:"article"`
}

func makePostArticleEndpoint(service article.UseCase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postArticleRequest)
		options := article.CreateArticleOptions{
			Title:       req.Article.Title,
			Description: req.Article.Description,
			Body:        req.Article.Body,
			TagList:     req.Article.TagList,
			Author:      userDomain.User{},
		}

		newArticle, err := service.Create(ctx, options)
		if err != nil {
			return nil, err
		}
		return postArticleResponse{Article: transformArticle(*newArticle)}, nil
	}
}
