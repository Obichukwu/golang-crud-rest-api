package main

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port                   string `mapstructure:"port"`
	MySqlConnectionString  string `mapstructure:"mysql_connection_string"`
	SqliteConnectionString string `mapstructure:"sqlite_connection_string"`
}

var AppConfig *Config

func LoadAppConfig() {
	log.Println("Loading Server Configurations...")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
