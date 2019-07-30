package main

import (
	"github.com/NavPool/navpool-api/config"
	"github.com/NavPool/navpool-api/controllers"
	"github.com/NavPool/navpool-api/database"
	"github.com/NavPool/navpool-api/middleware"
	"github.com/getsentry/raven-go"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	setReleaseMode()

	database.Migrate()

	if config.Get().Sentry.Active {
		_ = raven.SetDSN(config.Get().Sentry.DSN)
	}

	r := gin.New()

	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(gin.Recovery())
	r.Use(middleware.Authentication)
	r.Use(middleware.NetworkSelect)
	r.Use(middleware.ErrorHandler)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to NavPool Node API!")
	})

	userController := new(controllers.UserController)
	r.POST("/account", userController.Create)
	r.GET("/account/:id", userController.Read)

	infoController := new(controllers.InfoController)
	r.GET("/info", infoController.GetInfo)
	r.GET("/info/staking", infoController.GetStakingInfo)

	authenticated := r.Group("/auth/:apikey")

	addressController := new(controllers.AddressController)
	authenticated.POST("/address/:spendingAddress", addressController.Create)
	authenticated.GET("/address/:address/add/:signature", addressController.GetPoolAddress)
	authenticated.GET("/address/:address/validate", addressController.GetValidateAddress)

	communityFundController := new(controllers.CommunityFundController)
	authenticated.GET("/community-fund/proposal/list/:address", communityFundController.GetProposalVotes)
	authenticated.GET("/community-fund/payment-request/list/:address", communityFundController.GetPaymentRequestVotes)
	authenticated.POST("/community-fund/proposal/vote", communityFundController.PostProposalVote)
	authenticated.POST("/community-fund/payment-request/vote", communityFundController.PostPaymentRequestVote)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resource Not Found"})
	})

	_ = r.Run(":" + config.Get().Server.Port)
}

func setReleaseMode() {
	if config.Get().Debug == false {
		log.Printf("Mode: %s", gin.ReleaseMode)
		gin.SetMode(gin.ReleaseMode)
	} else {
		log.Printf("Mode: %s", gin.DebugMode)
		gin.SetMode(gin.DebugMode)
	}
}
