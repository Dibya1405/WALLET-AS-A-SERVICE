package entity

import "time"

type Wallet struct {
	// gorm.Model
	Id         string    `json:"walletId,omitempty" gorm:"primaryKey"`
	User       User      `json:"-"`
	UserId     string    `json:"userId,omitempty"`
	WalletName string    `json:"wallet_name,omitempty"`
	Balance    int       `json:"balance,omitempty"`
	IsBlocked  bool      `json:"is_blocked,omitempty"`
	IsDeleted  bool      `json:"is_deleted,omitempty"`
	Creation   time.Time `json:"created_at,omitempty"`
	Updation   time.Time `json:"updated_at,omitempty"`
}
