package controllers

import (
	"errors"
	"github.com/NavPool/navpool-api/helpers"
	"github.com/NavPool/navpool-api/navcoind"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InfoController struct{}

func (controller *InfoController) GetInfo(c *gin.Context) {
	getInfo, err := navcoind.Info{}.GetInfo()
	if err != nil {
		helpers.HandleError(c, ErrCouldNotGetInfo, http.StatusBadRequest)
		return
	}

	c.JSON(200, getInfo)
}

func (controller *InfoController) GetStakingInfo(c *gin.Context) {
	getStakingInfo, err := navcoind.Info{}.GetStakingInfo()
	if err != nil {
		helpers.HandleError(c, ErrCouldNotGetStakingInfo, http.StatusBadRequest)
		return
	}

	c.JSON(200, getStakingInfo)
}

var (
	ErrCouldNotGetInfo        = errors.New("Unable to retrieve info from node")
	ErrCouldNotGetStakingInfo = errors.New("Unable to retrieve staking info from node")
)
