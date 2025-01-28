package auth_test

import (
	"bytes"
	"encoding/json"
	"go/email-verification/configs"
	"go/email-verification/internal/auth"
	"go/email-verification/internal/user"
	"go/email-verification/pkg/db"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func bootstrap() (*auth.AuthHandler, sqlmock.Sqlmock, error) {
	database, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}))
	if err != nil {
		return nil, nil, err
	}

	userRepo := user.NewUserRepository(&db.Db{
		DB: gormDb,
	})

	handler := auth.AuthHandler{
		Config: &configs.Config{
			Auth: configs.AuthConfig{
				Secret: "secret",
			},
		},
		AuthService: auth.NewAuthService(userRepo),
	}

	return &handler, mock, nil
}

func TestLoginHandlerSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	if err != nil {
		t.Fatal(err)
	}

	rows := sqlmock.NewRows([]string{"email", "password"}).
		AddRow("altynbek2@altynbek.com", "$2a$10$jSZ5wo4.ogX3PyNVFJnu7uoWe7IZVzQjjuQTi0C61HwtR3iwdJAZy")

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "altynbek2@altynbek.com",
		Password: "1234",
	})

	reader := bytes.NewReader(data)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/auth/login", reader)
	handler.Login()(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}

func TestRegisterHandlerSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	if err != nil {
		t.Fatal(err)
	}

	rows := sqlmock.NewRows([]string{"email", "password", "name"})

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	data, _ := json.Marshal(&auth.RegistrationRequest{
		Email:    "altynbek3@altynbek.com",
		Password: "1234",
		Name:     "Altynbek",
	})

	reader := bytes.NewReader(data)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/auth/register", reader)
	handler.Register()(w, r)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status code %d, but got %d", http.StatusCreated, w.Code)
	}
}
