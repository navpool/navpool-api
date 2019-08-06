package resource

import (
	"errors"
	"github.com/NavPool/navpool-api/app/container"
	model "github.com/NavPool/navpool-api/app/model/user"
	"github.com/gin-gonic/gin"
)

type UserResource struct{}

func (r *UserResource) Create(c *gin.Context) {
	user, err := model.UserRepository().CreateUser(container.Container.Account)
	if err != nil {
		_ = c.Error(ErrUnableToCreateUser).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(200, user)
}

func (r *UserResource) Read(c *gin.Context) {
	user, err := model.UserRepository().GetByToken(container.Container.Account, c.GetString("token"))
	if err != nil {
		_ = c.Error(ErrUnableToGetUser).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(200, user)
}

var ErrUnableToGetUser = errors.New("Unable to get user")
var ErrUnableToCreateUser = errors.New("Unable to create user")
