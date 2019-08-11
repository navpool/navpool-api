package model_community_fund

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type VoteType string
type VoteChoice string

const (
	VoteTypeProposal       VoteType = "proposal"
	VoteTypePaymentRequest VoteType = "payment-request"
)

type Vote struct {
	ID             uint       `gorm:"primary_key" json:"-"`
	UserID         uuid.UUID  `gorm:"type:uuid;column:user_id;not null;" json:"-"`
	Type           VoteType   `json:"type"`
	StakingAddress string     `json:"staking_address"`
	Hash           string     `json:"hash"`
	Choice         string     `json:"vote"`
	CreatedAt      *time.Time `json:"_"`
	UpdatedAt      *time.Time `json:"_"`
	Committed      bool       `json:"committed"`
}
