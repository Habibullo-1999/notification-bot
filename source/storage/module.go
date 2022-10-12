package storage

import (
	"github.com/Habibullo-1999/notification-bot/source/entity"
	"github.com/Habibullo-1999/notification-bot/source/logger"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Module ...
var Module = fx.Options(
	fx.Provide(New),
)

// Param ...
type Param struct {
	fx.In
	DB     *gorm.DB
	Logger *logger.Logger
}

type storage struct {
	db     *gorm.DB
	logger *logger.Logger
}

type Storage interface {
	GetAllReservationLastFIveMinutes() ([]*entity.ReservationMeetingRoom, error)
	AddSendMessageTg(sendMessage *entity.SendMessageTg) error
	GetTelegramUserByUsername(username string) (*entity.TelegramUsers, error)
	AddTgUser(tgUser *entity.TelegramUsers) error
	GetUserByID(id string) (*entity.User, error)
	GetUserByTgUsername(username string) (*entity.User, error)
	GetAllSendMessagesByActive() ([]*entity.SendMessageTg, error)
	UpdateStatusDeActiveSendMessageByID(id string) error
	DeleteSendMessageByID(id string) error
}

func New(p Param) Storage {
	s := &storage{
		logger: p.Logger,
		db:     p.DB,
	}
	return s
}
