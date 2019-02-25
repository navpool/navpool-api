package main

import (
	"github.com/NavPool/navpool-api/config"
	"github.com/getsentry/raven-go"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sentry"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func init() {
	if config.Get().Sentry.Active == true {
		dsn := config.Get().Sentry.DSN
		log.Println("Sentry logging to ", dsn)
		raven.SetDSN(dsn)
	}
}

func main() {
	r := setupRouter()

	if config.Get().Ssl == false {
		r.Run(":" + config.Get().Server.Port)
	} else {
		log.Fatal(autotls.Run(r, config.Get().Server.Domain))
	}

	if config.Get().Sentry.Active == true {
		r.Use(sentry.Recovery(raven.DefaultClient, false))
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(networkSelect)
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(errorHandler)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to NavPool API!")
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resource Not Found"})
	})

	return r
}

func networkSelect(c *gin.Context) {
	switch network := c.GetHeader("Network"); network {
	case "testnet":
		config.Get().SelectedNetwork = network
		break
	case "mainnet":
		config.Get().SelectedNetwork = network
		break
	default:
		config.Get().SelectedNetwork = "mainnet"
	}

	c.Header("X-Network", config.Get().SelectedNetwork)
}

func errorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) == 0 {
		return
	}

	c.AbortWithStatusJSON(http.StatusBadRequest, c.Errors)
}
