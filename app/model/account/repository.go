package model_account

import (
	"github.com/NavPool/navpool-api/app/database"
	"github.com/jinzhu/gorm"
)

type repo struct {
	DB *gorm.DB
}

func AccountRepository() *repo {
	return &repo{
		DB: database.GetConnection(),
	}
}

func (p *repo) Create(username string, secret string, active bool) (*Account, error) {
	var account = &Account{Username: username, Secret: secret, Active: active}
	err := database.GetConnection().Create(account).Error

	return account, err
}

func (p *repo) GetByUsername(username string) (*Account, error) {
	var account = new(Account)
	err := database.GetConnection().Where(&Account{Username: username}).First(&account).Error

	return account, err
}
