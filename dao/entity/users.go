package entity

import "time"

type User struct {
	// gorm.Model - added id, created_at and updated_at and deleted_at
	Id        string    `json:"userId,omitempty" gorm:"primaryKey"`
	Name      string    `json:"full_name,omitempty"`
	Email     string    `json:"email_id,omitempty"`
	Mobile    int       `json:"mobile,omitempty"`
	Aadhaar   string    `json:"aadhaar_no,omitempty"`
	Creation  time.Time `json:"created_at,omitempty"`
	Updation  time.Time `json:"updated_at,omitempty"`
	Isdeleted bool      `json:"is_deleted,omitempty"`
}
