package auth

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegistrationResponse struct {
	Token string `json:"token"`
}

type RegistrationRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}
