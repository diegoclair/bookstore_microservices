package config

import (
	"github.com/diegoclair/microservice_oauth/domain/entity"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

// GetDBConfig to read initial config
func GetDBConfig() (config entity.InitialConfig) {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	config.Username = cast.ToString(viper.Get("DB_USER"))
	config.Password = cast.ToString(viper.Get("DB_PASSWORD"))
	config.Host = cast.ToString(viper.Get("DB_HOST"))
	config.Port = cast.ToString(viper.Get("DB_PORT"))
	config.DBName = cast.ToString(viper.Get("DB_NAME"))
	config.DBDefault = cast.ToString(viper.Get("DB_DEFAULT"))

	return
}
