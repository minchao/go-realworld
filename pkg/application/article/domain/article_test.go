package domain

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestArticle_CreateArticle(t *testing.T) {
	c := qt.New(t)

	article := Article{}

	c.Assert(article.TagList, qt.DeepEquals, []Tag(nil))
}

func TestArticle_AddTag(t *testing.T) {
	c := qt.New(t)
	article := Article{}

	article.AddTag("foo")

	c.Assert(article.TagList, qt.DeepEquals, []Tag{"foo"})
}
