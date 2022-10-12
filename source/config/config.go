package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"strings"
)

// Config ...
type Config interface {
	Get(key string) interface{}
	GetString(key string) string
	GetInt(key string) int64
	IsSet(key string) bool
	GetFloat(key string) float64
}

type config struct {
	cfg *viper.Viper
}

// New ...
func New() Config {
	cfg := viper.New()
	cfg.SetConfigName(".env")
	cfg.SetConfigType("env")
	cfg.AddConfigPath("./")
	cfg.AddConfigPath("../")
	cfg.AddConfigPath("../../../")

	if err := cfg.ReadInConfig(); err != nil {
		panic(fmt.Errorf("config - NewConfig() - ReadInConfig: %w", err))
	}

	cfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return &config{cfg: cfg}
}

// Module ...
var Module = fx.Provide(New)

func (c *config) Get(key string) interface{} {
	return c.cfg.Get(key)
}

func (c *config) GetString(key string) string {
	return c.cfg.GetString(key)
}

func (c *config) IsSet(key string) bool {
	return c.cfg.IsSet(key)
}

func (c *config) GetInt(key string) int64 {
	return c.cfg.GetInt64(key)
}

func (c *config) GetFloat(key string) float64 {
	return c.cfg.GetFloat64(key)
}
