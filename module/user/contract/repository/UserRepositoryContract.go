package contract

import (
	"context"
	"go-clean-architecture/module/user/model"
)

type UserRepository interface {
	Create(c context.Context, user *model.User) error
	Fetch(c context.Context) ([]model.User, error)
	GetByUsernameOrEmail(c context.Context, usernameOrEmail string) (model.User, error)
	GetByEmail(c context.Context, email string) (model.User, error)
	GetByID(c context.Context, id string) (model.User, error)
}
