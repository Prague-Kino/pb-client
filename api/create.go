package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Prague-Kino/cast"
	"github.com/Prague-Kino/pb-client/env"
	"github.com/Prague-Kino/pb-client/internal/errors"
	pb "github.com/Prague-Kino/pb-client/internal/pocketbase"
)

func (p *PocketBase) CreateKino(kino *cast.Kino) error {
	body, err := json.Marshal(kino)
	if err != nil {
		return err
	}

	return createRecord(pb.KinoCollection, p.authToken, body)
}

func (p *PocketBase) CreateFilm(kino *cast.Film) error {
	body, err := json.Marshal(kino)
	if err != nil {
		return err
	}

	return createRecord(pb.FilmCollection, p.authToken, body)
}

func (p *PocketBase) CreateScreening(screening *cast.Screening) error {
	body, err := json.Marshal(screening)
	if err != nil {
		return err
	}

	return createRecord(pb.ScreeningCollection, p.authToken, body)
}

// --------------------------------------

func createRecord(collection, token string, body []byte) error {
	envVars := env.GetEnvVars()
	postUrl := getPostUrl(envVars.BaseURL, collection)
	req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &errors.RecordCreateResponseReadError{Err: err}
	}

	if envVars.IsDev() {
		log.Printf("Status: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
		log.Printf("Response: %s", string(respBody))
	}

	if resp.StatusCode != http.StatusOK {
		return errors.PBErrorFromResponse(resp.StatusCode, respBody)
	}

	return nil
}

func getPostUrl(baseURL, collection string) string {
	return baseURL + "/api/collections/" + collection + "/records"
}
