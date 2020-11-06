package resource

import (
	"errors"
	"github.com/NavPool/navpool-api/internal/framework/error"
	"github.com/NavPool/navpool-api/internal/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type InfoResource struct {
	service *service.InfoService
}

func NewInfoResource(service *service.InfoService) *InfoResource {
	return &InfoResource{service}
}

func (r *InfoResource) GetInfo(c *gin.Context) {
	info, err := r.service.GetInfo()
	if err != nil {
		log.WithError(err).Error("Failed to getinfo")
		error.HandleError(c, ErrCouldNotGetInfo, http.StatusBadRequest)
		return
	}

	c.JSON(200, info)
}

func (r *InfoResource) GetStakingInfo(c *gin.Context) {
	stakingInfo, err := r.service.GetStakingInfo()
	if err != nil {
		log.WithError(err).Error("Failed to get staking info")
		error.HandleError(c, ErrCouldNotGetStakingInfo, http.StatusBadRequest)
		return
	}

	c.JSON(200, stakingInfo)
}

var (
	ErrCouldNotGetInfo        = errors.New("Unable to retrieve info from node")
	ErrCouldNotGetStakingInfo = errors.New("Unable to retrieve staking info from node")
)
