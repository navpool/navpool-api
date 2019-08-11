package resource

import (
	"errors"
	"github.com/NavPool/navpool-api/app/helpers"
	"github.com/NavPool/navpool-api/app/services/navcoin"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MessageResource struct{}

func (r *MessageResource) Validate(c *gin.Context) {
	var validateDto validateDto

	err := c.BindJSON(&validateDto)
	if err != nil {
		helpers.HandleError(c, ErrMessageRequestNotValid, http.StatusBadRequest)
		return
	}

	valid, err := navcoin.NewNavcoin().VerifyMessage(validateDto.Address, validateDto.Signature, validateDto.Message)
	if err != nil {
		helpers.HandleError(c, ErrUnableToValidate, http.StatusBadGateway)
		return
	}

	c.JSON(200, valid)
}

type validateDto struct {
	Address   string `json:"address" binding:"required"`
	Signature string `json:"signature" binding:"required"`
	Message   string `json:"message" binding:"required"`
}

var (
	ErrMessageRequestNotValid = errors.New("The vote is not valid")
	ErrUnableToValidate       = errors.New("The vote is not valid")
)
