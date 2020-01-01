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

type CreateArticleOptions struct {
	Title       string
	Description string
	Body        string
	TagList     []string
	Author      userDomain.User
}

func (s *Service) Create(ctx context.Context, options CreateArticleOptions) (*domain.Article, error) {
	now := time.Now()
	newArticle := &domain.Article{
		Slug:           slug.Make(options.Title),
		Title:          options.Title,
		Description:    options.Description,
		Body:           options.Body,
		TagList:        []domain.Tag{},
		CreatedAt:      now,
		UpdatedAt:      now,
		Favorited:      false,
		FavoritesCount: 0,
		Author:         options.Author,
	}
	for _, v := range options.TagList {
		tag := domain.Tag(v)
		newArticle.AddTag(tag)
		_ = s.tagService.Create(ctx, tag)
	}

	if err := s.articleRepository.Save(ctx, newArticle); err != nil {
		return nil, err
	}
	return newArticle, nil
}

func (s *Service) Get(ctx context.Context, slug string) (*domain.Article, error) {
	return s.articleRepository.Find(ctx, slug)
}

func (s *Service) GetAll(ctx context.Context, criteria domain.ArticleCriteria) ([]*domain.Article, error) {
	return s.articleRepository.FindBy(ctx, criteria)
}

func (s *Service) Update(ctx context.Context, article *domain.Article) error {
	return s.articleRepository.Save(ctx, article)
}

func (s *Service) Delete(ctx context.Context, slug string) error {
	return s.articleRepository.Delete(ctx, slug)
}
