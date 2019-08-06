package resource

import (
	"errors"
	"github.com/NavPool/navpool-api/app/container"
	"github.com/NavPool/navpool-api/app/helpers"
	addressModel "github.com/NavPool/navpool-api/app/model/address"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddressResource struct{}

func (r *AddressResource) Create(c *gin.Context) {
	c.JSON(200, nil)
}

func (r *AddressResource) Read(c *gin.Context) {
	address, err := addressModel.GetAddressBySpendingAddress(container.Container.User.ID, c.Param("spendingAddress"))
	if err != nil {
		helpers.HandleError(c, ErrAddressNotFound, http.StatusNotFound)
		return
	}

	c.JSON(200, address)
}

func (r *AddressResource) ReadAll(c *gin.Context) {
	addresses, err := addressModel.GetAddressesByUserId(container.Container.User.ID)
	if err != nil {
		helpers.HandleError(c, ErrAddressNotFound, http.StatusNotFound)
		return
	}

	c.JSON(200, addresses)
}

var ErrAddressNotFound = errors.New("Address not found")
