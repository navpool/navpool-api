package model_address

import (
	"github.com/NavPool/navpool-api/app/database"
	model "github.com/NavPool/navpool-api/app/model/user"
	"github.com/jinzhu/gorm"
)

type repo struct {
	DB *gorm.DB
}

func AddressRepository() *repo {
	return &repo{
		DB: database.GetConnection(),
	}
}

func (r *repo) GetAddressBySpendingAddress(user model.User, spendingAddress string) (*Address, error) {
	var address = new(Address)
	err := r.DB.Where(&Address{UserID: user.ID, SpendingAddress: spendingAddress}).First(&address).Error

	return address, err
}

func (r *repo) GetAddressByStakingAddress(user model.User, stakingAddress string) (*Address, error) {
	var address = new(Address)
	err := r.DB.Where(&Address{UserID: user.ID, StakingAddress: stakingAddress}).First(&address).Error

	return address, err
}

func (r *repo) GetAddressesByUser(user model.User) (*Address, error) {
	var address = new(Address)
	err := r.DB.Where(&Address{UserID: user.ID}).First(&address).Error

	return address, err
}
