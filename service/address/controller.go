package address

import (
	"errors"
	"github.com/NavPool/navpool-api/error"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct{}

func (controller *Controller) GetPoolAddress(c *gin.Context) {
	spendingAddress := c.Param("address")
	if spendingAddress == "" || !ValidateAddress(spendingAddress) {
		error.HandleError(c, ErrAddressNotValid, http.StatusBadRequest)
		return
	}

	poolAddress, err := GetPoolAddress(spendingAddress)
	if err != nil {
		error.HandleError(c, ErrPoolAddressNotAvailable, http.StatusInternalServerError)
		return
	}

	c.JSON(200, poolAddress)
}

var (
	ErrAddressNotValid         = errors.New("address not valid")
	ErrPoolAddressNotAvailable = errors.New("pool address not available")
)
