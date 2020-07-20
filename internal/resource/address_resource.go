package resource

import (
	b64 "encoding/base64"
	"github.com/NavPool/navpool-api/error"
	"github.com/NavPool/navpool-api/internal/config"
	"github.com/NavPool/navpool-api/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AddressResource struct {
	addressService *service.AddressService
}

func NewAddressResource(service *service.AddressService) *AddressResource {
	return &AddressResource{service}
}

func (r *AddressResource) GetPoolAddress(c *gin.Context) {
	spendingAddress := c.Param("address")
	validateAddress, err := r.addressService.ValidateAddress(spendingAddress)
	if err != nil {
		log.Println(err)
		error.HandleError(c, service.ErrUnableToValidateAddress, http.StatusInternalServerError)
		return
	}

	if validateAddress.Mine && config.Get().Debug == false {
		error.HandleError(c, service.ErrAddressIsMine, http.StatusBadRequest)
		return
	}

	if !validateAddress.Valid {
		error.HandleError(c, service.ErrAddressNotValid, http.StatusBadRequest)
		return
	}

	if validateAddress.ColdStaking {
		error.HandleError(c, service.ErrAddressIsColdStaking, http.StatusBadRequest)
		return
	}

	signature, _ := b64.StdEncoding.DecodeString(c.Param("signature"))
	verified, err := r.addressService.VerifySignature(spendingAddress, string(signature), "REGISTER FOR NAVPOOL")
	if err != nil || verified == false {
		error.HandleError(c, service.ErrSignatureNotValid, http.StatusBadRequest)
		return
	}

	poolAddress, err := r.addressService.GetPoolAddress(spendingAddress)
	if err != nil {
		error.HandleError(c, service.ErrPoolAddressNotAvailable, http.StatusInternalServerError)
		return
	}

	c.JSON(200, poolAddress)
}

func (r *AddressResource) GetValidateAddress(c *gin.Context) {
	address := c.Param("address")

	validateAddress, err := r.addressService.ValidateAddress(address)
	if err != nil {
		error.HandleError(c, service.ErrAddressNotValid, http.StatusBadRequest)
		return
	}

	c.JSON(200, validateAddress)
}
