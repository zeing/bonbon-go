package lineservice

import (
	"bonbon-go/client/twitterclient"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/gin-gonic/gin"
	"github.com/google/go-querystring/query"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/rs/zerolog/log"
)

type defaultServices struct {
	lbc *linebot.Client
	tc  twitterclient.TwitterClient
}

type LineBotService interface {
	Handler(ctx *gin.Context) (*linebot.Event, *twitter.Tweet, error)
	ReplyToUser(event *linebot.Event, replyMessage string) error
}

func Init(LineBot *linebot.Client, Twitter twitterclient.TwitterClient) LineBotService {
	return &defaultServices{
		lbc: LineBot,
		tc:  Twitter,
	}
}

func (svc *defaultServices) Handler(ctx *gin.Context) (*linebot.Event, *twitter.Tweet, error) {
	events, err := svc.lbc.ParseRequest(ctx.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			return nil, nil, err
		} else {
			return nil, nil, err
		}
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				return svc.handleText(event, message)
			//case *linebot.StickerMessage:
			//	svc.handleSticker(event, message)
			case *linebot.LocationMessage:
				return svc.handleLocation(event, message)
			case *linebot.ImageMessage:
				return svc.handleImage(event, message)
			}
		}
	}
	return nil, nil, nil
}

func (svc *defaultServices) handleText(event *linebot.Event, message *linebot.TextMessage) (*linebot.Event, *twitter.Tweet, error) {
	tweet, err := svc.tc.Tweet(message.Text, nil)
	if err != nil {
		return nil, nil, err
	}
	return event, tweet, err
}

func (svc *defaultServices) handleSticker(event *linebot.Event, message *linebot.StickerMessage) {
	replyMessage := fmt.Sprintf(
		"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
	if _, err := svc.lbc.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
		log.Logger.Print(err)
	}
}

type Options struct {
	Query string `url:"query"`
}

func (svc *defaultServices) handleLocation(event *linebot.Event, message *linebot.LocationMessage) (*linebot.Event, *twitter.Tweet, error) {
	q := fmt.Sprintf(
		"%f,%f", message.Latitude, message.Longitude)
	opt := Options{q}
	v, _ := query.Values(opt)
	query := v.Encode()

	location := fmt.Sprintf(
		"https://www.google.com/maps/search/?api=1&%s", query)
	messages := fmt.Sprintf(
		"%s %s", message.Title, location)

	params := &twitter.StatusUpdateParams{
		Lat:  &message.Latitude,
		Long: &message.Longitude,
	}
	tweet, err := svc.tc.Tweet(messages, params)
	if err != nil {
		return nil, nil, err
	}
	return event, tweet, err
}

func (svc *defaultServices) handleImage(event *linebot.Event, message *linebot.ImageMessage) (*linebot.Event, *twitter.Tweet, error) {
	content, err := svc.lbc.GetMessageContent(message.ID).Do()
	if err != nil {
		return nil, nil, err
	}

	media, err := svc.tc.UploadMedia(message.ID, content.Content)
	if err != nil {
		log.Logger.Err(err).Msg("error to upload media")
		return nil, nil, err
	}

	params := &twitter.StatusUpdateParams{
		MediaIds: []int64{media.MediaId},
	}
	tweet, err := svc.tc.Tweet("", params)
	if err != nil {
		return nil, nil, err
	}

	return event, tweet, nil
}

func (svc *defaultServices) ReplyToUser(event *linebot.Event, replyMessage string) error {
	if _, err := svc.lbc.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
		log.Logger.Print(err)
	}
	return nil
}
