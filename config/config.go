package config

import (
	"github.com/spf13/viper"

	"github.com/alfianyulianto/go-room-managament/halpers"
)

type AppConfig struct {
	BaseUrl      string
	DatabaseHost string
	DatabaseName string
	DatabasePort string
	DatabaseUser string
	DatabasePass string
	SecretKey    string
}

var Cfg AppConfig

func init() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	halpers.IfPanicError(err)

	Cfg = AppConfig{
		BaseUrl:      viper.GetString("APP_BASE_URL"),
		DatabaseHost: viper.GetString("DATABASE_HOST"),
		DatabaseName: viper.GetString("DATABASE_NAME"),
		DatabasePort: viper.GetString("DATABASE_PORT"),
		DatabaseUser: viper.GetString("DATABASE_USER"),
		DatabasePass: viper.GetString("DATABASE_PASS"),
		SecretKey:    viper.GetString("SECRET_KEY"),
	}
}
