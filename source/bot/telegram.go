package bot

import (
	"errors"
	"github.com/Habibullo-1999/notification-bot/source/config"
	"github.com/Habibullo-1999/notification-bot/source/entity"
	"github.com/Habibullo-1999/notification-bot/source/logger"
	"github.com/Habibullo-1999/notification-bot/source/storage"
	"github.com/go-co-op/gocron"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"log"
	"time"
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
	Config  config.Config
	Logger  *logger.Logger
	Storage storage.Storage
}

// Telegram ...
type Telegram interface {
	TelegramCon()
}

type telegram struct {
	config  config.Config
	logger  *logger.Logger
	storage storage.Storage
	bot     *tgbotapi.BotAPI
}

// New ...
func New(p Param) Telegram {
	t := &telegram{
		config:  p.Config,
		logger:  p.Logger,
		storage: p.Storage,
	}
	t.Connection()
	t.RunWorker()
	go t.TelegramCon()

	return t
}

func (t *telegram) Connection() {
	var token = t.config.GetString("TELEGRAM_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		t.logger.Log(err.Error(), "telegram", "Connection", "NewBotAPI")
		panic(err)
	}
	t.bot = bot
}

func (t *telegram) TelegramCon() {

	log.Printf("Authorized on account %s", t.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := t.bot.GetUpdatesChan(u)

	for update := range updates {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		if update.Message.Command() == "start" { // If we got a message
			var tgUser entity.TelegramUsers

			_, err := t.storage.GetUserByTgUsername(update.Message.From.UserName)
			if err != nil {
				switch {
				case errors.Is(err, gorm.ErrRecordNotFound):
					msg.Text = t.config.GetString("WELCOME_TEXT_UNREGISTERED_USER")
					t.bot.Send(msg)
				default:
					t.logger.PrintError(err, "bot", "TelegramCon", "GetUserByTgUsername")
					msg.Text = "Сорри, что та не так с сервером повторите позже !!!"
					t.bot.Send(msg)
				}
				continue
			}

			tgUser.ID = int(update.Message.From.ID)
			tgUser.FirstName = update.Message.From.FirstName
			tgUser.LastName = update.Message.From.LastName
			tgUser.Username = update.Message.From.UserName
			tgUser.LanguageCode = update.Message.From.LanguageCode
			tgUser.IsBot = update.Message.From.IsBot
			err = t.storage.AddTgUser(&tgUser)
			if err != nil {
				t.logger.PrintError(err, "bot", "TelegramCon", "AddTgUser")
				continue
			}

			msg.Text = t.config.GetString("WELCOME_TEXT")
		} else {
			msg.Text = t.config.GetString("UNKNOWN_COMMAND")
		}
		t.bot.Send(msg)
	}
}

func (t *telegram) RunWorker() {
	w := gocron.NewScheduler(time.UTC)

	w.Every(1).Minute().Do(t.SendNotification)
	w.StartAsync()
}

func (t *telegram) SendNotification() {
	reserves, err := t.storage.GetAllSendMessagesByActive()
	if err != nil {
		t.logger.PrintError(err, "bot", "SendNotification", "GetAllSendMessagesByActive")
		panic(err)
	}
	for _, reserv := range reserves {
		message := tgbotapi.NewMessage(int64(reserv.TgUserId), reserv.Text)
		_, err = t.bot.Send(message)
		if err != nil {
			t.logger.PrintError(err, "bot", "SendNotification", "NewMessage")
			return
		}
		err = t.storage.UpdateStatusDeActiveSendMessageByID(reserv.ID)
		if err != nil {
			t.logger.PrintError(err, "bot", "SendNotification", "NewMessage")
			log.Println(err)
		}
	}
}
