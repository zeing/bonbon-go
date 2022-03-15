package twitterclient

import (
	"bonbon-go/config"
	"encoding/base64"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"net/http"
)

type TwitterImpl struct {
	TwitterHttpClient *http.Client
	TwitterClient     *twitter.Client
}

type TwitterClient interface {
	Tweet(text string) (*twitter.Tweet, error)
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

func (tc TwitterImpl) Tweet(text string) (*twitter.Tweet, error) {
	// Send a Tweet
	tweet, _, err := tc.TwitterClient.Statuses.Update(text, nil)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}

func (tc TwitterImpl) UploadMedia(img base64.Encoding) (string, error) {
	//_, err := tc.TwitterHttpClient.Post("test", "", img)
	//if err != nil {
	//	return "", err
	//}
	return "", nil
}
