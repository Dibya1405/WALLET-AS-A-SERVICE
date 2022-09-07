package entity

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Id              string    `json:"transactionId,omitempty" gorm:"primaryKey"`
	User            User      `json:"-"`
	UserId          string    `json:"userId,omitempty"`
	Wallet          Wallet    `json:"-"`
	WalletId        string    `json:"walletId,omitempty"`
	TransactionType string    `json:"transaction_type,omitempty"`
	Amount          int       `json:"amount,omitempty"`
	Creation        time.Time `json:"created_at,omitempty"`
	Updation        time.Time `json:"updated_at,omitempty"`
}
