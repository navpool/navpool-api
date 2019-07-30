package controllers

import (
	"errors"
	"github.com/NavPool/navpool-api/account"
	"github.com/NavPool/navpool-api/config"
	"github.com/NavPool/navpool-api/helpers"
	"github.com/gin-gonic/gin"
	uuidService "github.com/satori/go.uuid"
	"net/http"
)

type UserController struct{}

func (controller *UserController) Create(c *gin.Context) {
	user, err := account.CreateUser(config.Get().Account)
	if err != nil {
		helpers.HandleError(c, ErrUnableToCreateUser, http.StatusBadRequest)
		return
	}

	c.JSON(200, user)
}

func (controller *UserController) Read(c *gin.Context) {
	var user *account.User

	id := c.GetString("id")
	uuid, err := uuidService.FromString(id)

	if err != nil {
		user, err = account.GetUserByUuid(config.Get().Account, uuid)
	} else {
		user, err = account.GetUserByApiToken(config.Get().Account, id)
	}

	if err != nil {
		_ = c.Error(ErrUnableToGetUser).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(200, user)
}

var (
	ErrUnableToGetUser    = errors.New("Unable to get user")
	ErrUnableToCreateUser = errors.New("Unable to create user")
)
