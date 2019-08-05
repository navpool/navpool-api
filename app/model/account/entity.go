package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Username string    `gorm:"unique" json:"username"`
	Secret   string    `json:"_"`
	Active   bool      `json:"active,not null"`
}

func (account *Account) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}
