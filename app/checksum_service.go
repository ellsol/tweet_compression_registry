package app

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/go-jet/jet/v2/postgres"
	"tweet_compression_registry/.gen/postgres/tcr/model"
	. "tweet_compression_registry/.gen/postgres/tcr/table"
)

type ChecksumService struct {
	db *sql.DB
}

func NewChecksumService(db *sql.DB) ChecksumService {
	return ChecksumService{db}
}

func (s *ChecksumService) GetTweet(context context.Context, data *GetTweetDTO) (*GetTweetResponseDTO, error) {
	stmt := Tweet.SELECT(Tweet.Original, Tweet.ID).FROM(Tweet).WHERE(Tweet.Checksum.EQ(postgres.String(data.Checksum)))
	dest := []model.Tweet{}

	err := stmt.Query(s.db, &dest)

	if err != nil {
		return nil, err
	}

	if len(dest) == 0 {
		return nil, fmt.Errorf("no tweet found with checksum: `%s`", data.Checksum)
	}

	foundTweet := dest[0]

	return &GetTweetResponseDTO{
		Id:            foundTweet.ID.String(),
		Tweet_Content: foundTweet.Original,
	}, nil
}

func (s *ChecksumService) GetPaginatedTweets(context context.Context, data *PaginationDTO) (*PaginatedTweetsResponseDTO, error) {
	stmt := Tweet.SELECT(Tweet.Original, Tweet.ID).FROM(Tweet).OFFSET(int64(data.Offset)).LIMIT(int64(data.Limit))

	dest := []model.Tweet{}
	err := stmt.Query(s.db, &dest)

	if err != nil {
		return nil, err
	}

	var tweets []GetTweetResponseDTO

	for _, tweet_model := range dest {
		tweet := GetTweetResponseDTO{Id: tweet_model.ID.String(), Tweet_Content: tweet_model.Original}
		tweets = append(tweets, tweet)
	}

	return &PaginatedTweetsResponseDTO{
		Tweets:     tweets,
		Pagination: *data,
	}, nil

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
