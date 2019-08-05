package resource

import (
	"errors"
	model "github.com/NavPool/navpool-api/app/model/user"
	"github.com/NavPool/navpool-api/app/session"
	"github.com/gin-gonic/gin"
)

type UserResource struct{}

func (r *UserResource) Create(c *gin.Context) {
	user, err := model.UserRepository().CreateUser(session.Account)
	if err != nil {
		_ = c.Error(ErrUnableToCreateUser).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(200, user)
}

func (r *UserResource) Read(c *gin.Context) {
	user, err := model.UserRepository().GetByToken(session.Account, c.GetString("token"))
	if err != nil {
		_ = c.Error(ErrUnableToGetUser).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(200, user)
}

var ErrUnableToGetUser = errors.New("Unable to get user")
var ErrUnableToCreateUser = errors.New("Unable to create user")
