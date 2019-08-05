package addressModel

import (
	"errors"
	uuidService "github.com/satori/go.uuid"
)

var (
	ErrAccountNotFound = errors.New("Account not found")
)

func GetAddressBySpendingAddress(userId uuidService.UUID, spendingAddress string) (address *Address, err error) {
	//db, err := database.NewConnection()
	//if err != nil {
	//	err = ErrAccountNotFound
	//	return nil, err
	//}
	//defer database.Close(db)
	//
	//err = db.Where(&Address{UserID: userId, SpendingAddress: spendingAddress}).First(&address).Error

	return
}

func GetAddressesByUserId(userId uuidService.UUID) (address []*Address, err error) {
	//db, err := database.NewConnection()
	//if err != nil {
	//	err = ErrAccountNotFound
	//	return nil, err
	//}
	//defer database.Close(db)
	//
	//err = db.Where(&Address{UserID: userId}).First(&address).Error
	return
}
