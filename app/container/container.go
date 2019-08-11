package container

import (
	"github.com/NavPool/navpool-api/app/config"
	accountModel "github.com/NavPool/navpool-api/app/model/account"
	userModel "github.com/NavPool/navpool-api/app/model/user"
)

var Container container

type container struct {
	Account accountModel.Account
	User    userModel.User
	Network *config.Network
}
