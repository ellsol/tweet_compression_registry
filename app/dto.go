package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// swagger:model UploadTweetDTO
type UploadTweetDTO struct {
	Payload string `json:"payload"`
}

// swagger:model UploadTweetResponseDTO
type UploadTweetResponseDTO struct {
	Id       string `json:"id"`
	Checksum string `json:"checksum"`
}

func (a *UploadTweetDTO) ReadAndValidate(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(a)

	if err != nil {
		return fmt.Errorf("body is empty")
	}

	return a.Validate()
}

func (data *UploadTweetDTO) Validate() error {
	if data.Payload == "" {
		return fmt.Errorf("payload cannot be empty")
	}
	return nil
}
