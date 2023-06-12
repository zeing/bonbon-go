package telegramservice

import (
	"bonbon-go/client/twitterclient"
	"bonbon-go/constant"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
)

type defaultServices struct {
	tgc *tgbotapi.BotAPI
	tc  twitterclient.TwitterClient
}

type TelegramService interface {
	Handler() error
	ReplyToUser(update tgbotapi.Update, replyMessage string) error
}

func Init(TelegramClient *tgbotapi.BotAPI, Twitter twitterclient.TwitterClient) TelegramService {
	return &defaultServices{
		tgc: TelegramClient,
		tc:  Twitter,
	}
}

func (service *defaultServices) Handler() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := service.tgc.GetUpdatesChan(u)

	// Create a new MessageConfig. We don't have text yet,
	// so we leave it empty.

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] eingtext %s", update.Message.From.UserName, update.Message)

			params := &twitter.StatusUpdateParams{}

			if len(update.Message.Photo) > 0 {
				file, err := service.tgc.GetFileDirectURL(update.Message.Photo[0].FileID)
				println("photo", file)
				media, err := service.tc.UploadMediaBase64(string(file), constant.TweetImage)
				if err != nil {
					log.Logger.Err(err).Msg("error to upload media")
					return err
				}
				println("media", media.MediaId)
				params.MediaIds = []int64{media.MediaId}

			}
			println("params", params)

			tweet, err := service.tc.Tweet(update.Message.Caption, params)
			if err != nil {
				log.Error().Err(err).Msg("telegram can't tweet")
			}

			replyMessage := fmt.Sprintf(
				"Tweeted !! | See at https://twitter.com/bon2_official/status/%s", tweet.IDStr)
			err = service.ReplyToUser(update, replyMessage)
			if err != nil {
				log.Error().Err(err).Msg("Error => Twitter Tweet Error")
				return err
			}
		}
	}

	return nil
}

func (service *defaultServices) ReplyToUser(update tgbotapi.Update, replyMessage string) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, replyMessage)
	msg.ReplyToMessageID = update.Message.MessageID
	service.tgc.Send(msg)
	return nil
}
