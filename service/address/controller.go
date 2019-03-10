package address

import (
	"errors"
	"github.com/NavPool/navpool-api/config"
	"github.com/NavPool/navpool-api/error"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Controller struct{}

func (controller *Controller) GetPoolAddress(c *gin.Context) {
	spendingAddress := c.Param("address")
	validateAddress, err := ValidateAddress(spendingAddress)
	if err != nil {
		log.Println(err)
		error.HandleError(c, ErrUnableToValidateAddress, http.StatusInternalServerError)
		return
	}

	if validateAddress.Mine && config.Get().Debug == false {
		error.HandleError(c, ErrAddressIsMine, http.StatusBadRequest)
		return
	}

	if !validateAddress.Valid {
		error.HandleError(c, ErrAddressNotValid, http.StatusBadRequest)
		return
	}

	if validateAddress.ColdStaking {
		error.HandleError(c, ErrAddressIsColdStaking, http.StatusBadRequest)
		return
	}

	poolAddress, err := GetPoolAddress(spendingAddress)
	if err != nil {
		error.HandleError(c, ErrPoolAddressNotAvailable, http.StatusInternalServerError)
		return
	}

	c.JSON(200, poolAddress)
}

func (controller *Controller) GetValidateAddress(c *gin.Context) {
	address := c.Param("address")

	validateAddress, err := ValidateAddress(address)
	if err != nil {
		error.HandleError(c, ErrAddressNotValid, http.StatusBadRequest)
		return
	}

	c.JSON(200, validateAddress)
}

var (
	ErrUnableToValidateAddress = errors.New("unable to validate address")
	ErrAddressNotValid         = errors.New("address not valid")
	ErrAddressIsColdStaking    = errors.New("address is a cold staking address")
	ErrAddressIsMine           = errors.New("address is owned by the pool")
	ErrPoolAddressNotAvailable = errors.New("pool address not available")
)
