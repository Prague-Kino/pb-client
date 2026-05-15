package errors

import (
	"encoding/json"
	"fmt"
	"strings"
)

func PBErrorFromResponse(statusCode int, respBody []byte) error {
	var pbErr PocketBaseError

	err := json.Unmarshal(respBody, &pbErr)
	if err != nil {
		return fmt.Errorf(
			"failed to parse pocketbase error response: %w",
			err,
		)
	}

	return &pbErr
}

type PocketBaseError struct {
	Data    map[string]PBFieldError `json:"data"`
	Message string                  `json:"message"`
	Status  int                     `json:"status"`
}

type PBFieldError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e PocketBaseError) Error() string {
	var details []string

	for field, fieldErr := range e.Data {
		details = append(
			details,
			fmt.Sprintf(
				"%s: %s (error code: %s)", field, fieldErr.Message, fieldErr.Code,
			),
		)
	}

	message := fmt.Sprintf(
		"PocketBase Error %d: %s",
		e.Status,
		e.Message,
	)

	if len(details) > 0 {
		message = fmt.Sprintf(
			"%s\n \t> %s\n",
			message,
			strings.Join(details, "\n\t> "),
		)
	}

	return message
}

type AuthResponseReadError struct {
	Err error
}

func (e AuthResponseReadError) Error() string {
	return fmt.Sprintf("Failed to read Auth response body: %s", e.Err)
}

type AuthResponseParseError struct {
	Err error
}

func (e AuthResponseParseError) Error() string {
	return fmt.Sprintf("Failed to parse Auth response body: %s", e.Err)
}

type RecordCreateResponseReadError struct {
	Err error
}

func (e RecordCreateResponseReadError) Error() string {
	return fmt.Sprintf("Failed to read record creation response body: %s", e.Err)
}
