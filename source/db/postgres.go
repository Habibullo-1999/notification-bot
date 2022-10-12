package db

import (
	"fmt"
	"github.com/Habibullo-1999/notification-bot/source/config"
	"github.com/Habibullo-1999/notification-bot/source/logger"
	"go.uber.org/fx"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"

	l "gorm.io/gorm/logger"
)

// Module ...
var Module = fx.Options(
	fx.Provide(NewConn),
)

//Param is all dependencies needed to create database connection
type Param struct {
	fx.In
	Config config.Config
	Logger *logger.Logger
}

type configuration struct {
	User       string
	Password   string
	Network    string
	Address    string
	Port       string
	DbName     string
	Parameters string
}

func (cfg configuration) formatDSN() string {
	//return cfg.User + ":" + cfg.Password + "@" + cfg.Network + "(" + cfg.Address + ")/" + cfg.DbName + "?" + cfg.Parameters
	return "host=" + cfg.Address + " user=" + cfg.User + " password=" + cfg.Password + " dbname=" + cfg.DbName + " port=" + cfg.Port + " sslmode=disable TimeZone=Asia/Dushanbe"
}

//NewConn is provider for type *Pool
func NewConn(p Param) (db *gorm.DB, err error) {
	cfg := configuration{
		User:       p.Config.GetString("db_username"),
		Password:   p.Config.GetString("db_password"),
		Network:    p.Config.GetString("db_protocol"),
		Address:    p.Config.GetString("db_address"),
		Port:       p.Config.GetString("db_port"),
		DbName:     p.Config.GetString("db_name"),
		Parameters: p.Config.GetString("db_params"),
	}

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "log/database.log",
		MaxSize:    10,
		MaxBackups: 10,
		MaxAge:     10,
	})

	newLogger := l.New(
		log.New(w, "\r\n", log.LstdFlags), // io writer
		l.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      l.Info,      // Log level
			Colorful:      false,       // Disable color
		},
	)

	URL := cfg.formatDSN()
	db, err = gorm.Open(postgres.Open(URL), &gorm.Config{Logger: newLogger})
	if err != nil {
		return nil, fmt.Errorf("db - NewConn() - gormOpen: %w", err)
	}

	return db, nil
}
