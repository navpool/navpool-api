package info

import (
	"errors"
	"github.com/NavPool/navpool-api/error"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct{}

var (
	ErrCouldNotGetInfo        = errors.New("Unable to retrieve info from node")
	ErrCouldNotGetStakingInfo = errors.New("Unable to retrieve staking info from node")
)

func (controller *Controller) GetInfo(c *gin.Context) {
	info, err := GetInfo()
	if err != nil {
		error.HandleError(c, ErrCouldNotGetInfo, http.StatusBadRequest)
		return
	}

	c.JSON(200, info)
}

func (controller *Controller) GetStakingInfo(c *gin.Context) {
	stakingInfo, err := GetStakingInfo()
	if err != nil {
		error.HandleError(c, ErrCouldNotGetStakingInfo, http.StatusBadRequest)
		return
	}

	c.JSON(200, stakingInfo)
}
