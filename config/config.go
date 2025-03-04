package config

import (
	"log"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type Server struct {
	Port int
}

type Db struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
	Sslmode  string
}

type Config struct {
	Server *Server
	Db     *Db
}

var (
	once           sync.Once
	configInstance *Config
)

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("Config err", err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			log.Fatal("Config reading error", err)
		}
	})
	return configInstance
}
