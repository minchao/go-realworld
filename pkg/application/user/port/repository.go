package port

import (
	"context"

	"github.com/minchao/go-realworld/pkg/application/user/domain"
)

type UserRepository interface {
	Find(ctx context.Context, email string) (*domain.User, error)
}
