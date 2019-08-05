package resource

import (
	"errors"
	"github.com/NavPool/navpool-api/app/helpers"
	addressModel "github.com/NavPool/navpool-api/app/model/address"
	userModel "github.com/NavPool/navpool-api/app/model/user"
	"github.com/NavPool/navpool-api/app/session"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddressResource struct{}

func (r *AddressResource) Create(c *gin.Context) {
	user, err := userModel.UserRepository().GetByToken(session.Account, c.Param("token"))
	if err != nil {
		helpers.HandleError(c, ErrUnableToGetUser, http.StatusBadRequest)
		return
	}

	c.JSON(200, user)
}

func (r *AddressResource) Read(c *gin.Context) {
	user, err := userModel.UserRepository().GetByToken(session.Account, c.Param("token"))
	if err != nil {
		helpers.HandleError(c, ErrUnableToGetUser, http.StatusBadRequest)
		return
	}

	address, err := addressModel.GetAddressBySpendingAddress(user.ID, c.Param("spendingAddress"))
	if err != nil {
		helpers.HandleError(c, ErrAddressNotFound, http.StatusNotFound)
		return
	}

	c.JSON(200, address)
}

func (r *AddressResource) ReadAll(c *gin.Context) {
	user, err := userModel.UserRepository().GetByToken(session.Account, c.Param("token"))
	if err != nil {
		helpers.HandleError(c, ErrUnableToGetUser, http.StatusBadRequest)
		return
	}

	addresses, err := addressModel.GetAddressesByUserId(user.ID)
	if err != nil {
		helpers.HandleError(c, ErrAddressNotFound, http.StatusNotFound)
		return
	}

	c.JSON(200, addresses)
}

var ErrAddressNotFound = errors.New("Address not found")
