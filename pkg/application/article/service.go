package article

import (
	"context"
	"time"

	"github.com/gosimple/slug"

	"github.com/minchao/go-realworld/pkg/application/article/domain"
	"github.com/minchao/go-realworld/pkg/application/article/port"
	userDomain "github.com/minchao/go-realworld/pkg/application/user/domain"
)

type Service struct {
	articleRepository port.ArticleRepository
	tagService        port.CreateTagService
}

func NewArticleService(articleRepo port.ArticleRepository, tagService port.CreateTagService) *Service {
	return &Service{
		articleRepository: articleRepo,
		tagService:        tagService,
	}
}

type CreationData struct {
	Title       string
	Description string
	Body        string
	TagList     []string
	Author      userDomain.User
}

func (s *Service) Create(ctx context.Context, data CreationData) (*domain.Article, error) {
	now := time.Now()
	newArticle := &domain.Article{
		Slug:           slug.Make(data.Title),
		Title:          data.Title,
		Description:    data.Description,
		Body:           data.Body,
		TagList:        []domain.Tag{},
		CreatedAt:      now,
		UpdatedAt:      now,
		Favorited:      false,
		FavoritesCount: 0,
		Author:         data.Author,
	}
	for _, v := range data.TagList {
		tag := domain.Tag(v)
		newArticle.AddTag(tag)
		_ = s.tagService.Create(ctx, tag)
	}

	return newArticle, s.articleRepository.Save(ctx, newArticle)
}

func (s *Service) Get(ctx context.Context, slug string) (*domain.Article, error) {
	return s.articleRepository.Find(ctx, slug)
}

func (s *Service) GetAll(ctx context.Context, criteria domain.ArticleCriteria) ([]*domain.Article, error) {
	return s.articleRepository.FindBy(ctx, criteria)
}

type MutationData struct {
	Title       *string
	Description *string
	Body        *string
}

func (s *Service) Update(ctx context.Context, slug string, data MutationData) (*domain.Article, error) {
	article, err := s.articleRepository.Find(ctx, slug)
	if err != nil {
		return nil, err
	}

	if data.Title != nil {
		article.Title = *data.Title
	}
	if data.Description != nil {
		article.Description = *data.Description
	}
	if data.Body != nil {
		article.Body = *data.Body
	}

	article.UpdatedAt = time.Now()

	return article, s.articleRepository.Save(ctx, article)
}

func (s *Service) Delete(ctx context.Context, slug string) error {
	return s.articleRepository.Delete(ctx, slug)
}
