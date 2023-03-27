package app

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

// TODO List tweets paginated    GET /tweet?offset=10&limit=10
// TODO Get specific tweet     GET /tweet/bychecksum/{checksum}

type Controller struct {
	checksumService *ChecksumService
}

func NewController(service *ChecksumService) Controller {
	return Controller{
		service,
	}
}

func (c Controller) InsertTweet() func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/", c.UploadTweet)
	}
}

func (c Controller) RetrieveTweet() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/{checksum}", c.GetTweet)
	}
}

// swagger:parameters uploadTweet
type uploadTweet struct {
	// in:body
	Body UploadTweetDTO
}

// swagger:route POST /tweet tweet uploadTweet
// Uploads a tweet
//
// Responses:
//
//	200: UploadTweetResponseDTO
func (c Controller) UploadTweet(w http.ResponseWriter, r *http.Request) {
	data := &UploadTweetDTO{}
	if err := data.ReadAndValidate(r); err != nil {
		RespondWithJSON(w, BadRequest(err.Error()))
		return
	}

	result, err := c.checksumService.NewTweet(r.Context(), data)

	if err != nil {
		RespondWithJSON(w, InternalError(err.Error()))
		return
	}

	RespondWithJSON(w, OK(result))
}

func (c Controller) GetTweet(w http.ResponseWriter, r *http.Request) {
	data := &GetTweetDTO{}
	if err := data.ReadAndValidate(r); err != nil {
		RespondWithJSON(w, BadRequest(err.Error()))
		return
	}

	result, err := c.checksumService.GetTweet(r.Context(), data)

	if err != nil {
		RespondWithJSON(w, InternalError(err.Error()))
		return
	}

	RespondWithJSON(w, OK(result))
}
