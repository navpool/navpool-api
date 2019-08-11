package model_user

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;" json:"id"`
	Account   uuid.UUID  `json:"account"`
	Token     string     `gorm:"unique" json:"token,omitempty"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"update_at,omitempty"`
	Active    bool       `json:"active,not null"`
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}
