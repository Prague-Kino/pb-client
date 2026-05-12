package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Prague-Kino/pb-client/env"
	"github.com/Prague-Kino/pb-client/internal/errors"
	pb "github.com/Prague-Kino/pb-client/internal/pocketbase"
	"github.com/Prague-Kino/pb-client/models"
)

type Payload map[string]string

func GetAuthToken() (string, error) {
	envVars := env.GetEnvVars()

	payload := createAuthRequestPayload(envVars)
	body, _ := json.Marshal(payload)

	fullURL := getAuthWithPasswordEndpoint(envVars.BaseURL)

	resp, err := http.Post(fullURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", &errors.AuthResponseReadError{Err: err}
	}

	if envVars.IsDev() {
		log.Printf("Auth status: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
		log.Printf("Auth response: %s", string(respBody))
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.PBErrorFromResponse(resp.StatusCode, respBody)
	}

	var authResp models.AuthResponse
	if err := json.Unmarshal(respBody, &authResp); err != nil {
		return "", &errors.AuthResponseParseError{Err: err}
	}

	return authResp.Token, nil
}

func getAuthWithPasswordEndpoint(baseURL string) string {
	return baseURL + "/api/collections/" + pb.AuthCollection + "/auth-with-password"
}

func createAuthRequestPayload(envVars *models.EnvVars) Payload {
	return Payload{
		"identity": envVars.AdminEmail,
		"password": envVars.AdminPassword,
	}
}
