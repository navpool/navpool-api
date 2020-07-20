package config

import (
	"errors"
	"github.com/NavPool/navpool-api/internal/framework/param"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

var (
	ErrNetworkNotFound = errors.New("Network not found")
)

type Config struct {
	Debug     bool
	Ssl       bool
	Signature bool
	Server    ServerConfig
	Network   map[string]Network
	Sentry    Sentry
}

type ServerConfig struct {
	Port int
}

type Network struct {
	Name     string
	Host     string
	Port     int
	Username string
	Password string
}

type Sentry struct {
	Active bool
	DSN    string
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.WithError(err).Fatal("Unable to init config")
	}
}

func (c *Config) ActiveNetwork() (*Network, error) {
	for _, n := range c.Network {
		if n.Name == param.GetGlobalParam("network", "").(string) {
			return &n, nil
		}
	}

	return nil, ErrNetworkNotFound
}

func Get() *Config {
	networks := make(map[string]Network)

	networks["mainnet"] = Network{
		Name:     "mainnet",
		Host:     getString("NAVCOIND_MAINNET_HOST", "localhost"),
		Port:     getInt("NAVCOIND_MAINNET_PORT", 8332),
		Username: getString("NAVCOIND_MAINNET_USERNAME", "user"),
		Password: getString("NAVCOIND_MAINNET_PASSWORD", "password"),
	}

	networks["testnet"] = Network{
		Name:     "testnet",
		Host:     getString("NAVCOIND_TESTNET_HOST", "localhost"),
		Port:     getInt("NAVCOIND_TESTNET_PORT", 8332),
		Username: getString("NAVCOIND_TESTNET_USERNAME", "user"),
		Password: getString("NAVCOIND_TESTNET_PASSWORD", "password"),
	}

	return &Config{
		Debug:     getBool("DEBUG", true),
		Ssl:       getBool("SSL", true),
		Signature: getBool("SIGNATURE", true),
		Server: ServerConfig{
			Port: getInt("SERVER_PORT", 8080),
		},
		Network: networks,
		Sentry: Sentry{
			Active: getBool("SENTRY_ACTIVE", false),
			DSN:    getString("SENTRY_DSN", ""),
		},
	}
}

func getString(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func getInt(key string, defaultValue int) int {
	valStr := getString(key, "")
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}

	return defaultValue
}

func getUint(key string, defaultValue uint) uint {
	return uint(getInt(key, int(defaultValue)))
}

func getUint64(key string, defaultValue uint) uint64 {
	return uint64(getInt(key, int(defaultValue)))
}

func getBool(key string, defaultValue bool) bool {
	valStr := getString(key, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultValue
}

func getSlice(key string, defaultVal []string, sep string) []string {
	valStr := getString(key, "")
	if valStr == "" {
		return defaultVal
	}

	return strings.Split(valStr, sep)
}
