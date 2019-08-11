package config

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

type Config struct {
	Env       string
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

	DB DbConfig `yaml:"db"`

	Networks []Network `yaml:"networks"`
}

type DbConfig struct {
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
		_, b, _, _ := runtime.Caller(0)
		configPath := filepath.Dir(b)
		viper.AddConfigPath(configPath + "/../..")

		env := getEnv()
		viper.SetConfigName("config." + env)

		instance = &Config{Env: env}
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
		}
		log.Printf("Using config: %s", viper.ConfigFileUsed())

		if err := viper.Unmarshal(instance); err != nil {
			log.Fatal(err)
		}
	})

	return instance
}

func getEnv() string {
	env := "prod"

	if flag.Lookup("test.v") != nil {
		env = "test"
	} else if len(os.Args) > 1 {
		env = os.Args[1]
	}

	return env
}
