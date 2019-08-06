package routes

import (
	"github.com/NavPool/navpool-api/app/middleware"
	resource "github.com/NavPool/navpool-api/app/resources"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes() *gin.Engine {
	r := gin.Default()

	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(gin.Recovery())
	r.Use(middleware.NetworkSelect)
	r.Use(middleware.ErrorHandler)

	addressResource := new(resource.AddressResource)
	communityFundResource := new(resource.CommunityFundResource)
	infoResource := new(resource.InfoResource)
	userResource := new(resource.UserResource)

	public := r.Group("")
	public.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to NavPool Node API!")
	})

	protected := r.Group("")
	//protected.Use(authorization.Authorization)

	protected.GET("/info", infoResource.GetInfo)
	protected.GET("/info/staking", infoResource.GetStakingInfo)

	private := protected.Group("")
	private.Use(middleware.UserToken)

	private.POST("/user", userResource.Create)
	private.GET("/user/:token", userResource.Read)

	private.POST("/address/:spendingAddress", addressResource.Create)
	private.GET("/address", addressResource.ReadAll)
	private.GET("/address/:spendingAddress", addressResource.Read)

	private.GET("/community-fund/proposal/list/:address", communityFundResource.GetProposalVotes)
	private.GET("/community-fund/payment-request/list/:address", communityFundResource.GetPaymentRequestVotes)
	private.POST("/community-fund/proposal/vote", communityFundResource.PostProposalVote)
	private.POST("/community-fund/payment-request/vote", communityFundResource.PostPaymentRequestVote)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resource Not Found"})
	})

	return r
}
