package account

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/NavPool/navpool-api/database"
	"github.com/NavPool/navpool-api/helpers"
	uuidService "github.com/satori/go.uuid"
	"time"
)

var (
	ErrUserNotFound = errors.New("User not found")
)

func GetUserByUuid(account string, uuid uuidService.UUID) (user *User, err error) {
	db, err := database.NewConnection()
	if err != nil {
		err = ErrUserNotFound
		return
	}
	defer database.Close(db)

	db.Where(&User{Account: account, ID: uuid}).First(&user)

	return
}

func GetUserByApiToken(account string, apiKey string) (user *User, err error) {
	db, err := database.NewConnection()
	if err != nil {
		err = ErrUserNotFound
		return
	}
	defer database.Close(db)

	db.Where(&User{Account: account, ApiKey: apiKey}).First(&user)

	return
}

func CreateUser(account string) (user *User, err error) {
	db, err := database.NewConnection()
	if err != nil {
		helpers.LogError(err)
		return
	}
	defer database.Close(db)

	user = &User{Account: account, ApiKey: tokenGenerator(), Active: true}
	err = db.Create(user).Error
	if err != nil {
		helpers.LogError(err)
		return
	}

	return
}

func tokenGenerator() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

type User struct {
	ID          uuidService.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Account     string           `json:"_"`
	ApiKey      string           `gorm:"unique" json:"api_key,omitempty"`
	LastLoginAt *time.Time       `json:"last_login_at,omitempty"`
	DeletedAt   *time.Time       `sql:"index" json:"deleted_at,omitempty"`
	CreatedAt   *time.Time       `json:"created_at,omitempty"`
	UpdatedAt   *time.Time       `json:"update_at,omitempty"`
	Active      bool             `json:"active,not null"`
}
