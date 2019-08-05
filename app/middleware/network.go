package middleware

import (
	"github.com/NavPool/navpool-api/app/session"
	"github.com/gin-gonic/gin"
)

func NetworkSelect(c *gin.Context) {
	switch network := c.GetHeader("Network"); network {
	case "testnet":
	case "mainnet":
		session.Network = network
		break
	default:
		session.Network = "mainnet"
	}

	c.Header("X-Network", session.Network)
}
