package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"sync"
)

type Config struct {
	Debug     bool
	Ssl       bool
	Signature bool

	Sentry struct {
		Active bool
		DSN    string
	}

	Server struct {
		Port   string
		Domain string
	}

	DB DBConfig `yaml:"db"`

	Networks        []Network
	SelectedNetwork string

	Account string
}

type DBConfig struct {
	Dialect  string `yaml:"dialect"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbName"`
	SSLMode  string `yaml:"sslMode"`
	LogMode  bool   `yaml:"logMode"`
}

type Network struct {
	Name     string
	Host     string
	Port     int
	Username string
	Password string
}

var instance *Config
var once sync.Once

func Get() *Config {
	once.Do(func() {
		log.Println("Creating Config")
		var env = "prod"
		if len(os.Args) > 1 {
			env = os.Args[1]
		}

		viper.SetConfigName("config." + env)
		viper.AddConfigPath(".")

		instance = &Config{}

		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
		}

		if err := viper.Unmarshal(instance); err != nil {
			log.Fatal(err)
		}
	})

	return instance
}
