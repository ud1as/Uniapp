package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
	Timeout  string
	Addr     string
}

func NewAppConfig() (*Config, error) {
	// path: path=./heml/developers/config.json
	viper.SetConfigType("json") // Look for specific type
	filepath := "./heml/developers/config.json"
	if os.Getenv("path") != "" {
		filepath = os.Getenv("path")
	}
	file, err := os.Open(filepath)
	if err != nil {
		panic("failed to open file with config")
	}
	err = viper.ReadConfig(file)
	if err != nil {
		panic("failed to read config")
	}

	conf := &Config{
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.name_db"),
		viper.GetString("db.Timeout"),
		viper.GetString("Addr"),
	}
	return conf, nil
}
