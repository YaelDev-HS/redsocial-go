package data

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  []byte    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	Enabled   bool      `json:"enabled"`
}

func (*User) TableName() string {
	return "users"
}

type UserModel struct {
	db *gorm.DB
}

func (m *UserModel) Create(user *User) error {
	return m.db.Create(user).Error
}
