package controllers

import (
	b64 "encoding/base64"
	"errors"
	"github.com/NavPool/navpool-api/account"
	"github.com/NavPool/navpool-api/config"
	"github.com/NavPool/navpool-api/helpers"
	"github.com/NavPool/navpool-api/navcoind"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddressController struct{}

func (controller *AddressController) Create(c *gin.Context) {
	user, err := account.GetUserByApiToken(config.Get().Account, c.Param("apiKey"))
	if err != nil {
		helpers.HandleError(c, ErrUnableToGetUser, http.StatusBadRequest)
		return
	}

	c.JSON(200, nil)
}

func (controller *AddressController) GetPoolAddress(c *gin.Context) {
	addressService := navcoind.Address{}

	spendingAddress := c.Param("address")
	validateAddress, err := addressService.GetValidateAddress(spendingAddress)
	if err != nil {
		helpers.HandleError(c, ErrUnableToValidateAddress, http.StatusInternalServerError)
		return
	}

	if !validateAddress.Valid {
		helpers.HandleError(c, ErrAddressNotValid, http.StatusBadRequest)
		return
	}

	if validateAddress.ColdStaking {
		helpers.HandleError(c, ErrAddressIsColdStaking, http.StatusBadRequest)
		return
	}

	signature, _ := b64.StdEncoding.DecodeString(c.Param("signature"))
	verified, err := navcoind.Signature{}.VerifySignature(spendingAddress, string(signature), "REGISTER FOR NAVPOOL")
	if err != nil || verified == false {
		helpers.HandleError(c, ErrSignatureNotValid, http.StatusBadRequest)
		return
	}

	poolAddress, err := addressService.GetPoolAddress(spendingAddress)
	if err != nil {
		helpers.HandleError(c, ErrPoolAddressNotAvailable, http.StatusInternalServerError)
		return
	}

	c.JSON(200, poolAddress)
}

func (controller *AddressController) GetValidateAddress(c *gin.Context) {
	validateAddress, err := navcoind.Address{}.GetValidateAddress(c.Param("address"))
	if err != nil {
		helpers.HandleError(c, ErrAddressNotValid, http.StatusBadRequest)
		return
	}

	c.JSON(200, validateAddress)
}

var (
	ErrUnableToValidateAddress = errors.New("unable to validate address")
	ErrAddressNotValid         = errors.New("address not valid")
	ErrAddressIsColdStaking    = errors.New("address is a cold staking address")
	ErrPoolAddressNotAvailable = errors.New("pool address not available")
	ErrSignatureNotValid       = errors.New("signature not valid")
)
