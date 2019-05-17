package main

import (
	"github.com/NavPool/navpool-api/config"
	"github.com/NavPool/navpool-api/service/address"
	"github.com/NavPool/navpool-api/service/communityFund"
	"github.com/getsentry/raven-go"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	if config.Get().Debug == false {
		gin.SetMode(gin.ReleaseMode)
	}

	if config.Get().Sentry.Active {
		raven.SetDSN(config.Get().Sentry.DSN)
	}

	r := gin.New()

	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(gin.Recovery())
	r.Use(networkSelect)
	r.Use(errorHandler)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to NavPool Node API!")
	})

	addressController := new(address.Controller)
	r.GET("/address/:address/add/:signature", addressController.GetPoolAddress)
	r.GET("/address/:address/validate", addressController.GetValidateAddress)

	communityFundController := new(communityFund.Controller)
	r.GET("/community-fund/:type/list/:address", communityFundController.GetVotes)
	r.POST("/community-fund/:type/vote", communityFundController.PostVote)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resource Not Found"})
	})

	_ = r.Run(":" + config.Get().Server.Port)
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
