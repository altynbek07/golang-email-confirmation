package auth

import (
	"go/email-verification/configs"
	"go/email-verification/pkg/jwt"
	"go/email-verification/pkg/mail"
	"go/email-verification/pkg/middleware"
	"go/email-verification/pkg/req"
	"go/email-verification/pkg/res"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
	*mail.EmailService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
	*mail.EmailService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:       deps.Config,
		AuthService:  deps.AuthService,
		EmailService: deps.EmailService,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
	router.Handle("GET /auth/test", middleware.IsAuthed(handler.Test(), deps.Config))
	// router.HandleFunc("POST /email/verification-notification", handler.SendVerification())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](w, r)

		if err != nil {
			return
		}

		email, err := handler.AuthService.Login(body.Email, body.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.JWTData{Email: email})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := LoginResponse{
			Token: token,
		}
		res.Json(w, data, http.StatusOK)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegistrationRequest](w, r)

		if err != nil {
			return
		}

		email, err := handler.AuthService.Register(body.Email, body.Password, body.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.JWTData{Email: email})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := RegistrationResponse{
			Token: token,
		}
		res.Json(w, data, http.StatusCreated)
	}
}

func (handler *AuthHandler) Test() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler.EmailService.Send("a@a.kz", "Test subject", "Test body")
		res.Json(w, "email sended", http.StatusOK)
	}
}

// func (handler *AuthHandler) SendVerification() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		body, err := req.HandleBody[RegistrationRequest](w, r)

// 		if err != nil {
// 			return
// 		}

// 		email, err := handler.AuthService.Register(body.Email, body.Password, body.Name)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusUnauthorized)
// 			return
// 		}

// 		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.JWTData{Email: email})
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		data := RegistrationResponse{
// 			Token: token,
// 		}
// 		res.Json(w, data, 201)
// 	}
// }
