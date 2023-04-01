package app

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
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

// swagger:model GetTweetResponseDTO
type GetTweetResponseDTO struct {
	Id            string `json:"id"`
	Tweet_Content string `json:"tweet"`
}

type PaginationDTO struct {
	Page   int `json:"page"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Size   int `json:"size"`
}

type PaginatedTweetsResponseDTO struct {
	Tweets     []GetTweetResponseDTO `json:"tweets"`
	Pagination PaginationDTO         `json:"pagination"`
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
		return fmt.Errorf("no checksum provided")
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
