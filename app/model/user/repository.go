package model

import (
	"crypto/rand"
	"fmt"
	"github.com/NavPool/navpool-api/app/database"
	"github.com/NavPool/navpool-api/app/model/account"
	"github.com/jinzhu/gorm"
)

type repo struct {
	DB *gorm.DB
}

func UserRepository() *repo {
	return &repo{
		DB: database.GetConnection(),
	}
}

func (r *repo) CreateUser(account model.Account) (*User, error) {
	var user = &User{Account: account.ID, Token: generateToken()}
	err := r.DB.Create(&user).Error

	return user, err
}

func (r *repo) GetByToken(account model.Account, token string) (*User, error) {
	var user = new(User)
	err := r.DB.Where(&User{Account: account.ID, Token: token}).First(&user).Error

	return user, err
}

func generateToken() string {
	b := make([]byte, 16)
	rand.Read(b)

	return fmt.Sprintf("%x", b)
}
