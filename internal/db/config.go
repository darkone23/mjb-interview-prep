package db

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type Sqlite3Conf struct {
	ConnectionUrl string `mapstructure:"dbname"`
}

type DbConf struct {
	Sqlite Sqlite3Conf `mapstructure:"sqlite3"`
}

func LoadConfig() DbConf {
	v := viper.New()
	config_path := os.Getenv("SQL_CONFIG")

	if config_path == "" {
		log.Fatalln("SQL_CONFIG env var required!")
	}

	v.SetConfigFile(config_path)

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("couldn't load db config: %s\n", err)
	}

	var c DbConf
	if err := v.Unmarshal(&c); err != nil {
		log.Printf("couldn't read config: %s\n", err)
	}

	if c.Sqlite.ConnectionUrl == "" {
		log.Fatalf("Connection URL should not be empty!")
	}

	return c

}
