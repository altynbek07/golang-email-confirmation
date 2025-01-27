package auth_test

import (
	"go/email-confirmation/internal/auth"
	"go/email-confirmation/internal/user"
	"testing"
)

type MockUserRepository struct{}

func (repo *MockUserRepository) Create(u *user.User) (*user.User, error) {
	return &user.User{
		Email: "a@a.kz",
	}, nil
}

func (repo *MockUserRepository) FindByEmail(email string) (*user.User, error) {
	return nil, nil
}

func TestRegisterSuccess(t *testing.T) {
	const initialEmail = "a@a.kz"
	authService := auth.NewAuthService(&MockUserRepository{})
	email, err := authService.Register(initialEmail, "1234", "Altynbek")
	if err != nil {
		t.Fatal(err)
	}

	if email != initialEmail {
		t.Fatalf("Expected email %s, but got %s", initialEmail, email)
	}
}
