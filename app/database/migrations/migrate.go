package migrations

import (
	"github.com/NavPool/navpool-api/app/config"
	"github.com/NavPool/navpool-api/app/database"
	model "github.com/NavPool/navpool-api/app/model/account"
	"os/user"
)

func Migrate() {
	database.GetConnection().AutoMigrate(
		&model.Account{},
		&user.User{})

	if config.Get().Env == "dev" {
		model.AccountRepository().Create("admin", "admin", true)
	}
}
