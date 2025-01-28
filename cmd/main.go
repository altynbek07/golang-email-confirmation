package main

import (
	"fmt"
	"go/email-verification/configs"
	"go/email-verification/internal/auth"
	"go/email-verification/internal/user"
	"go/email-verification/pkg/db"
	"go/email-verification/pkg/mail"
	"net/http"
)

func App() *http.ServeMux {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	// Repositories
	userRepository := user.NewUserRepository(db)

	// Services
	authService := auth.NewAuthService(userRepository)
	mailService := mail.NewEmailService(conf)

	// Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:       conf,
		AuthService:  authService,
		EmailService: mailService,
	})

	return router
}

func main() {
	app := App()

	server := http.Server{
		Addr:    ":8081",
		Handler: app,
	}
	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
