package entity

import "time"

type User struct {
	// gorm.Model - added id, created_at and updated_at and deleted_at
	Id        string    `json:"userId,omitempty" gorm:"primaryKey"`
	Name      string    `json:"full_name,omitempty" gorm:"NOT NULL"`
	Email     string    `json:"email_id,omitempty" gorm:"NOT NULL"`
	Mobile    int       `json:"mobile,omitempty" gorm:"NOT NULL"`
	Aadhaar   int       `json:"aadhaar_no,omitempty" gorm:"NOT NULL"`
	Creation  time.Time `json:"created_at,omitempty"`
	Updation  time.Time `json:"updated_at,omitempty"`
	IsDeleted bool      `json:"is_deleted,omitempty"`
}
