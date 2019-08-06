package resource

import (
	"github.com/gin-gonic/gin"
)

type MessageResource struct{}

func (r *MessageResource) Validate(c *gin.Context) {
	c.JSON(200, nil)
}
