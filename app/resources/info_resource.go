package resource

import (
	"errors"
	"github.com/NavPool/navpool-api/app/helpers"
	"github.com/NavPool/navpool-api/app/services/navcoind"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InfoResource struct{}

func (r *InfoResource) GetInfo(c *gin.Context) {
	getInfo, err := navcoind.Info{}.GetInfo()
	if err != nil {
		helpers.HandleError(c, ErrCouldNotGetInfo, http.StatusBadRequest)
		return
	}

	c.JSON(200, getInfo)
}

func (r *InfoResource) GetStakingInfo(c *gin.Context) {
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
