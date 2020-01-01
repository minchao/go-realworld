package port

import (
	articlePort "github.com/minchao/go-realworld/pkg/application/article/port"
	userPort "github.com/minchao/go-realworld/pkg/application/user/port"
)

type Manager interface {
	Article() articlePort.ArticleRepository
	Tag() articlePort.TagRepository
	User() userPort.UserRepository
}
