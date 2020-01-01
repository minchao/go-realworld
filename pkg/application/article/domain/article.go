package domain

import (
	"errors"
	"time"

	"github.com/minchao/go-realworld/pkg/application/user/domain"
)

var (
	ErrArticleNotFound = errors.New("article not found")
)

type Article struct {
	Slug           string
	Title          string
	Description    string
	Body           string
	TagList        []Tag
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Favorited      bool
	FavoritesCount int
	Author         domain.User
}

func (a *Article) AddTag(tag Tag) {
	a.TagList = append(a.TagList, tag)
}

type ArticleCriteria struct {
	Tag       string
	Author    string
	Favorited string
	Limit     uint
	Offset    uint
}

type Comment struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Body      string
	Author    domain.User
}
