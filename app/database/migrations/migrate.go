package migrations

import (
	"github.com/NavPool/navpool-api/app/config"
	"github.com/NavPool/navpool-api/app/database"
	model_account "github.com/NavPool/navpool-api/app/model/account"
	model_community_fund "github.com/NavPool/navpool-api/app/model/community_fund"
	model_user "github.com/NavPool/navpool-api/app/model/user"
)

func Migrate() {
	database.GetConnection().AutoMigrate(
		&model_account.Account{},
		&model_user.User{},
		&model_community_fund.Vote{},
	)

	if config.Get().Env == "dev" {
		model_account.AccountRepository().Create("admin", "admin", true)
	}
}
