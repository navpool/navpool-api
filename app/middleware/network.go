package middleware

import (
	"errors"
	"github.com/NavPool/navpool-api/app/config"
	"github.com/NavPool/navpool-api/app/container"
	"github.com/NavPool/navpool-api/app/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NetworkSelect(c *gin.Context) {
	network := c.GetHeader("Network")
	if network == "" {
		network = "mainnet"
	}

	networks := config.Get().Networks
	for i, _ := range networks {
		if networks[i].Name == network {
			container.Container.Network = &networks[i]
			c.Header("X-Network", container.Container.Network.Name)
			return
		}
	}

	helpers.HandleError(c, ErrNetworkNotFound, http.StatusBadRequest)
}

var (
	ErrNetworkNotFound = errors.New("Network not found")
)
