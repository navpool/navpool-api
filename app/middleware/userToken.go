package middleware

import (
	"errors"
	"github.com/NavPool/navpool-api/app/helpers"
	"github.com/NavPool/navpool-api/app/model/user"
	"github.com/NavPool/navpool-api/app/session"
	"github.com/gin-gonic/gin"
	"log"
)

func UserToken(c *gin.Context) {
	token := c.GetHeader("api-token")
	account := session.Account

	if len(token) != 0 {
		user, err := model.UserRepository().GetByToken(account, token)
		if err != nil {
			log.Printf("Invalid token($%s) for account (%s)", token, account.ID)
			helpers.HandleError(c, ErrInvalidUserToken, 401)
			return
		}

		session.User = *user
	}
}

var (
	ErrInvalidUserToken = errors.New("Invalid user token")
)
