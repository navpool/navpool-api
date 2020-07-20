package main

import (
	"fmt"
	"github.com/NavPool/navpool-api/generated/dic"
	"github.com/NavPool/navpool-api/internal/config"
	"github.com/NavPool/navpool-api/internal/framework"
	"github.com/NavPool/navpool-api/internal/resource"
	"github.com/getsentry/raven-go"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sarulabs/dingo/v3"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var container *dic.Container

func main() {
	config.Init()
	container, _ = dic.NewContainer(dingo.App)

	if config.Get().Debug {
		log.SetLevel(log.DebugLevel)
	}

	framework.SetReleaseMode(config.Get().Debug)

	if config.Get().Sentry.Active {
		raven.SetDSN(config.Get().Sentry.DSN)
	}

	r := gin.New()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(framework.Cors())
	r.Use(framework.NetworkSelect)
	r.Use(framework.Options)
	r.Use(framework.ErrorHandler)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to NavPool Node API!")
	})

	addressResource := resource.NewAddressResource(container.GetAddressService())
	r.GET("/address/:address/add/:signature", addressResource.GetPoolAddress)
	r.GET("/address/:address/validate", addressResource.GetValidateAddress)

	daoResource := resource.NewDaoResource(container.GetDaoService())
	r.GET("/dao/:type/list/:address", daoResource.GetVotes)
	r.POST("/dao/:type/vote", daoResource.PostVote)

	// BC support for community-fund path
	r.GET("/community-fund/:type/list/:address", daoResource.GetVotes)
	r.POST("/community-fund/:type/vote", daoResource.PostVote)

	infoResource := resource.NewInfoResource(container.GetInfoService())
	r.GET("/info", infoResource.GetInfo)
	r.GET("/info/staking", infoResource.GetStakingInfo)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resource Not Found"})
	})

	_ = r.Run(fmt.Sprintf(":%d", config.Get().Server.Port))
}
