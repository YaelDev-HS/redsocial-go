package data

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Nickname  string    `json:"nickname"`
	Password  []byte    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	IsEnabled bool      `json:"enabled"`
}

func (*User) TableName() string {
	return "users"
}

type UserModel struct {
	db *gorm.DB
}

func (m *UserModel) Create(user *User) error {
	err := m.db.Create(user).Error

	if err != nil && strings.Contains(err.Error(), DuplicatedKey) {
		return ErrDuplicatedKey
	}

	return err
}
