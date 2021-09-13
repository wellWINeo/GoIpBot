package GoIpBot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramRoutingFunction func(msg  *tgbotapi.Message) (tgbotapi.Chattable, error)
type TelegramMessageHandler func(userID int64, tag string, args []string) (string, error)

type BotConfig struct {
	Token   string
	Debug   bool
	Offset  int
	Timeout int
}

type Bot struct {
	API    *tgbotapi.BotAPI
	Config BotConfig
	UpdateConfig tgbotapi.UpdateConfig
}

func NewBot(config BotConfig) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		return nil, err
	}

	bot.Debug = config.Debug

	u := tgbotapi.NewUpdate(config.Offset)
	u.Timeout = config.Timeout

	return &Bot{
		API:          bot,
		Config:       config,
		UpdateConfig: u,
	}, nil
}

func (b *Bot) Poll() (chan error, chan tgbotapi.Message, chan tgbotapi.MessageConfig) {
	errs := make(chan error)
	in := make(chan tgbotapi.Message)
	out := make(chan tgbotapi.MessageConfig)
	updates, err := b.API.GetUpdatesChan(b.UpdateConfig)
	if err != nil {
		errs <- err
	}

	go func () {
		for update := range updates {
			Log("bot.go").Info("message sent to queue")
			in <- *update.Message
		}
	}()

	go func () {
		for msg := range out {
			Log("bot.go").Info("message received from queue")
			b.API.Send(msg)
		}
	}()

	return errs, in, out
}
