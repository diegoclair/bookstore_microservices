package config

import (
	"github.com/diegoclair/microservice_user/domain/entity"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

// GetDBConfig to read initial config
func GetDBConfig() (config entity.InitialConfig) {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	config.Username = cast.ToString(viper.Get("mysql_users_username"))
	config.Password = cast.ToString(viper.Get("mysql_users_password"))
	config.Host = cast.ToString(viper.Get("mysql_users_host"))
	config.Schema = cast.ToString(viper.Get("mysql_users_schema"))

	return
}
