package bot

import (
	"fmt"
	"github.com/Habibullo-1999/notification-bot/source/config"
	"github.com/Habibullo-1999/notification-bot/source/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/fx"
	"log"
)

// Module ...
var Module = fx.Options(
	fx.Invoke(
		New,
	),
)

// Param ...
type Param struct {
	fx.In
	Config config.Config
	Logger *logger.Logger
}

// Telegram ...
type Telegram interface {
	TelegramCon()
}

type telegram struct {
	config config.Config
	logger *logger.Logger
}

// New ...
func New(p Param) Telegram {
	t := &telegram{
		config: p.Config,
		logger: p.Logger,
	}
	t.TelegramCon()

	return t
}

func (t *telegram) TelegramCon() {
	var token = t.config.GetString("TELEGRAM_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		t.logger.Log(err.Error(), "bot/telegram", "TelegramCon", "tgbotapi.NewBotAPI")
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	fmt.Println(bot.GetWebhookInfo())

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
