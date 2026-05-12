package env

import (
	"os"

	"github.com/Prague-Kino/pb-client/models"
	"github.com/joho/godotenv"
)

const (
	PocketBaseURL           = "POCKETBASE_URL"
	PocketBaseAdminEmail    = "POCKETBASE_ADMIN_EMAIL"
	PocketBaseAdminPassword = "POCKETBASE_ADMIN_PASS"
	AppEnvironment          = "APP_ENV"
)

func LoadEnvVars(filenames ...string) error {
	err := godotenv.Load(filenames...)
	return err
}

func GetEnvVars() *models.EnvVars {
	baseURL := os.Getenv(PocketBaseURL)
	email := os.Getenv(PocketBaseAdminEmail)
	password := os.Getenv(PocketBaseAdminPassword)
	environment := os.Getenv(AppEnvironment)

	return &models.EnvVars{
		BaseURL:        baseURL,
		AdminEmail:     email,
		AdminPassword:  password,
		AppEnvironment: environment,
	}
}
