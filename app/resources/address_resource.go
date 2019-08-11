package resource

import (
	"errors"
	"github.com/NavPool/navpool-api/app/container"
	"github.com/NavPool/navpool-api/app/helpers"
	addressModel "github.com/NavPool/navpool-api/app/model/address"
	"github.com/NavPool/navpool-api/app/services/navcoin"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddressResource struct{}

func (r *AddressResource) Create(c *gin.Context) {
	createAddressDto := new(CreateAddressDto)
	if err := c.BindJSON(&createAddressDto); err != nil {
		helpers.HandleError(c, navcoin.ErrUnableToCreateAddress, http.StatusBadRequest)
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
	}

	address := addressModel.Address{
		UserID: container.Container.User.ID,
	}

	err := addressModel.AddressRepository().DB.Save(address).Error
	if err != nil {
		helpers.HandleError(c, navcoin.ErrUnableToCreateAddress, http.StatusInternalServerError)
		return
	}

	poolAddress, err := navcoin.NewNavcoin().CreatePoolAddress(createAddressDto.SpendingAddress, address.ID)
	if err != nil {
		helpers.HandleError(c, navcoin.ErrUnableToCreateAddress, http.StatusInternalServerError)
		return
	}

	address.SpendingAddress = poolAddress.SpendingAddress
	address.StakingAddress = poolAddress.StakingAddress
	address.ColdStakingAddress = poolAddress.ColdStakingAddress

	err = addressModel.AddressRepository().DB.Save(address).Error
	if err != nil {
		addressModel.AddressRepository().DB.Delete(addressModel.Address{ID: address.ID})
		helpers.HandleError(c, navcoin.ErrUnableToCreateAddress, http.StatusInternalServerError)
		return
	}

	c.JSON(200, poolAddress)
}

func (r *AddressResource) Read(c *gin.Context) {
	address, err := addressModel.AddressRepository().GetAddressBySpendingAddress(
		container.Container.User,
		c.Param("spendingAddress"),
	)

	if err != nil {
		helpers.HandleError(c, ErrAddressNotFound, http.StatusNotFound)
		return
	}

	c.JSON(200, address)
}

func (r *AddressResource) ReadAll(c *gin.Context) {
	addresses, err := addressModel.AddressRepository().GetAddressesByUser(container.Container.User)
	if err != nil {
		helpers.HandleError(c, ErrAddressNotFound, http.StatusNotFound)
		return
	}

	c.JSON(200, addresses)
}

type CreateAddressDto struct {
	SpendingAddress string `json:"spendingaddress" binding:"required"`
}

var ErrAddressNotFound = errors.New("Address not found")
