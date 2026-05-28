package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/Prague-Kino/pb-client/internal/errors"
	"github.com/Prague-Kino/pb-client/models"
)

type PocketBase struct {
	authToken string
	baseURL   string
}

func NewPocketBase(baseURL string) (*PocketBase, error) {
	token, err := GetAuthToken(baseURL)
	if err != nil {
		return nil, err
	}

	return &PocketBase{
		authToken: token,
		baseURL:   baseURL,
	}, nil
}

func (pb *PocketBase) post(collection string, body []byte) error {
	postUrl := fmt.Sprintf("%s/collections/%s/records", pb.baseURL, collection)

	req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	if pb.authToken != "" {
		req.Header.Set("Authorization", pb.authToken)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	resBpdy, err := io.ReadAll(res.Body)
	if err != nil {
		return &errors.InvalidResponseError{Err: err, Url: postUrl}
	}

	if res.StatusCode != http.StatusOK {
		return errors.PBErrorFromResponse(res.StatusCode, resBpdy)
	}

	return nil
}

func (pb *PocketBase) get(collection string, filter *models.Filter, expand *models.Expand) ([]byte, error) {
	params := url.Values{}
	params.Set("perPage", "200")
	if filter != nil {
		params.Set("filter", filter.String())
	}
	if expand != nil {
		params.Set("expand", expand.String())
	}

	getUrl := fmt.Sprintf(
		"%s/api/collections/%s/records?%s",
		pb.baseURL,
		collection,
		params.Encode(),
	)
	fmt.Println(">>> ", getUrl)

	res, err := http.Get(getUrl)
	if err != nil {
		return nil, &errors.HttpGetError{
			Url: getUrl,
			Err: err,
		}
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, &errors.InvalidResponseError{
			Url: getUrl,
			Err: err,
		}
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.PBErrorFromResponse(res.StatusCode, body)
	}

	return body, nil
}
