package di

import (
	"go/email-verification/internal/user"
)

type IUserRepository interface {
	Create(user *user.User) (*user.User, error)
	FindByEmail(email string) (*user.User, error)
}
