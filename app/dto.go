package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// swagger:model UploadTweetDTO
type UploadTweetDTO struct {
	Payload string `json:"payload"`
}

type GetTweetDTO struct {
	Checksum string
}

// swagger:model UploadTweetResponseDTO
type UploadTweetResponseDTO struct {
	Id       string `json:"id"`
	Checksum string `json:"checksum"`
}

type GetTweetResponseDTO struct {
	Id            string `json:"id"`
	Tweet_Content string `json:"tweet"`
}

func (a *UploadTweetDTO) ReadAndValidate(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(a)

	if err != nil {
		return fmt.Errorf("body is empty")
	}

	return a.Validate()
}

func (a *GetTweetDTO) ReadAndValidate(r *http.Request) error {
	checksum_val := chi.URLParam(r, "checksum")

	if checksum_val == "" {
		return fmt.Errorf("body is empty")
	}

	a.Checksum = checksum_val

	return a.Validate()
}

func (data *UploadTweetDTO) Validate() error {
	if data.Payload == "" {
		return fmt.Errorf("payload cannot be empty")
	}
	return nil
}

func (data *GetTweetDTO) Validate() error {
	if data.Checksum == "" {
		return fmt.Errorf("checksum cannot be empty")
	}
	return nil
}
