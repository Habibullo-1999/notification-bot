package service

import (
	"github.com/Habibullo-1999/notification-bot/source/config"
	"github.com/Habibullo-1999/notification-bot/source/logger"
	"github.com/Habibullo-1999/notification-bot/source/storage"
	"github.com/go-co-op/gocron"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"time"
)

// Module ...
var Module = fx.Options(
	fx.Invoke(New),
)

// Param ...
type Param struct {
	fx.In
	DB      *gorm.DB
	Logger  *logger.Logger
	Storage storage.Storage
	Config  config.Config
}

type service struct {
	db      *gorm.DB
	logger  *logger.Logger
	storage storage.Storage
	config  config.Config
}

// Service ...
type Service interface {
	WorkerAddSendMessage()
}

// New ...
func New(p Param) Service {
	s := &service{
		logger:  p.Logger,
		db:      p.DB,
		storage: p.Storage,
		config:  p.Config,
	}
	s.RunWorker()
	return s
}

func (s service) RunWorker() {
	w := gocron.NewScheduler(time.UTC)

	w.Every(5).Minute().Do(s.WorkerAddSendMessage)
	w.StartAsync()
}
