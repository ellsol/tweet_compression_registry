package app

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"tweet_compression_registry/.gen/postgres/tcr/model"

	. "tweet_compression_registry/.gen/postgres/tcr/table"
)

type ChecksumService struct {
	db *sql.DB
}

func NewChecksumService(db *sql.DB) ChecksumService {
	return ChecksumService{db}
}

func (s *ChecksumService) NewTweet(context context.Context, data *UploadTweetDTO) (*UploadTweetResponseDTO, error) {
	c := sha256.Sum256([]byte(data.Payload))
	check := hex.EncodeToString(c[:])

	stmt := Tweet.INSERT(Tweet.Checksum, Tweet.Original).VALUES(check, data.Payload).RETURNING(Tweet.AllColumns)
	dest := []model.Tweet{}

	err := stmt.Query(s.db, &dest)

	if err != nil {
		return nil, err
	}

	if len(dest) == 0 {
		return nil, fmt.Errorf("no tweet created")
	}

	newTweet := dest[0]

	return &UploadTweetResponseDTO{
		Id:       newTweet.ID.String(),
		Checksum: newTweet.Checksum,
	}, nil
}
