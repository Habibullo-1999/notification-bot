package main

import (
	"github.com/Habibullo-1999/notification-bot/source/bot"
	"github.com/Habibullo-1999/notification-bot/source/config"
	"github.com/Habibullo-1999/notification-bot/source/db"
	"github.com/Habibullo-1999/notification-bot/source/logger"
	"github.com/Habibullo-1999/notification-bot/source/service"
	"github.com/Habibullo-1999/notification-bot/source/storage"
	"github.com/go-resty/resty/v2"
	"go.uber.org/fx"
	"os"
)

func main() {
	args := os.Args[1:]
	var arg string
	if len(args) > 0 {
		arg = args[0]
	}
	var mainModules = fx.Options(
		fx.Provide(resty.New),
		config.Module,
		logger.Module,
		db.Module,
		fx.Provide(func() string {
			return arg
		}),
		service.Module,
		storage.Module,
		bot.Module,
	)
	fx.New(mainModules).Run()
}
