package twitterclient

import (
	"bonbon-go/config"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/rs/zerolog/log"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type TwitterImpl struct {
	TwitterHttpClient *http.Client
	TwitterClient     *twitter.Client
}

type TwitterClient interface {
	Tweet(text string, params *twitter.StatusUpdateParams) (*twitter.Tweet, error)
	UploadMedia(filename string, media io.Reader, mediaType string) (*UploadMediaResponse, error)
}

func NewTwitterClient() TwitterClient {

	configTwitter := oauth1.NewConfig(config.App.Twitter.ConsumerKey, config.App.Twitter.ConsumerSecret)
	token := oauth1.NewToken(config.App.Twitter.AccessTokenKey, config.App.Twitter.AccessTokenSecret)

	httpClient := configTwitter.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	return &TwitterImpl{
		TwitterHttpClient: httpClient,
		TwitterClient:     client,
	}
}

func (tc TwitterImpl) Tweet(text string, params *twitter.StatusUpdateParams) (*twitter.Tweet, error) {
	// Send a Tweet
	tweet, _, err := tc.TwitterClient.Statuses.Update(text, params)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}

type UploadMediaRequest struct {
	Media []byte `json:"media"`
}

type UploadMediaResponse struct {
	MediaId       int64  `json:"media_id"`
	MediaIdString string `json:"media_id_string"`
}

func (tc TwitterImpl) UploadMedia(filename string, media io.Reader, mediaType string) (*UploadMediaResponse, error) {

	// create body form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// create media paramater
	fw, err := writer.CreateFormFile("media", filename)
	if err != nil {
		log.Error().Err(err).Msg("add field to form error")
		return nil, err
	}
	// copy to form
	_, err = io.Copy(fw, media)
	if err != nil {
		return nil, err
	}

	writer.WriteField("media_category", mediaType)

	// close form
	writer.Close()

	url := fmt.Sprintf(
		"https://upload.twitter.com/1.1/media/upload.json")
	res, err := tc.TwitterHttpClient.Post(url, writer.FormDataContentType(), bytes.NewReader(body.Bytes()))
	if err != nil {
		log.Error().Err(err).Msg("tc.TwitterHttpClient.Post error")
		return nil, err
	}

	resByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error().Err(err).Msg("read response error")
		return nil, err
	}

	response := &UploadMediaResponse{}
	err = json.Unmarshal(resByte, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
