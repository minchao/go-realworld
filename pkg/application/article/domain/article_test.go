package domain

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestArticle_AddTag(t *testing.T) {
	c := qt.New(t)

	article := Article{}
	c.Assert(article.TagList, qt.DeepEquals, []Tag(nil))

	article.AddTag("foo")
	c.Assert(article.TagList, qt.DeepEquals, []Tag{"foo"})
	article.AddTag("bar")
	c.Assert(article.TagList, qt.DeepEquals, []Tag{"foo", "bar"})
}
