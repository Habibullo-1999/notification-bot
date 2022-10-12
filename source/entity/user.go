package entity

import "time"

// User struct
type User struct {
	ID           string     `json:"id" gorm:"primary_key"`
	Name         string     `json:"name"`
	LastName     string     `json:"last_name"`
	UniqueID     string     `json:"uniqueID"`
	Email        string     `json:"email" gorm:"unique"`
	TgAccount    string     `json:"tg_account" gorm:"unique" `
	DepartmentID string     `json:"department_id"`
	City         string     `json:"city"`
	Password     string     `json:"password,omitempty"`
	Role         string     `json:"role,omitempty"`
	Active       bool       `json:"active"`
	CreatedAt    *time.Time `json:"-" gorm:"index"`
	UpdatedAt    *time.Time `json:"-" gorm:"index"`
}

// TelegramUsers struct
type TelegramUsers struct {
	ID           int
	FirstName    string
	LastName     string
	Username     string
	LanguageCode string
	IsBot        bool
}
