package jwt_test

import (
	"go/email-confirmation/pkg/jwt"
	"testing"
)

func TestJWTCreate(t *testing.T) {
	const email = "a@a.kz"

	jwtService := jwt.NewJWT("295e0e2f5de5c91bc973834b017512dcace479b6ae9fa635de03756ddd0433fd")

	token, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}

	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal("Token is not valid")
	}

	if data.Email != email {
		t.Fatalf("Expected email %s, but got %s", email, data.Email)
	}
}
