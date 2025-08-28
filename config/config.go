package config

import (
	"github.com/spf13/viper"

	"github.com/alfianyulianto/go-room-managament/halpers"
)

type AppConfig struct {
	BaseUrl string
}

var Cfg AppConfig

func init() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	halpers.IfPanicError(err)

	Cfg = AppConfig{
		BaseUrl: viper.GetString("APP_BASE_URL"),
	}
}
