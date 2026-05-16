package data

import (
	"math"
	"time"

	"gorm.io/gorm"
)

type Message struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"-"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsEnabled bool      `json:"is_enabled"`
	User      *User     `json:"user" gorm:"foreignKey:UserID"`
}

type MessageModel struct {
	db *gorm.DB
}

func (m *MessageModel) Create(message *Message) error {
	return m.db.Omit("created_at", "updated_at", "is_enabled").Create(message).Error
}

func (m *MessageModel) FindAll(limit, page int) ([]*Message, error) {
	var messages []*Message

	offset := math.Max(float64(limit*(page-1)), 0)
	err := m.db.Offset(int(offset)).Limit(limit).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "nickname")
		}).
		Find(&messages).Error

	if err != nil {
		return nil, err
	}

	return messages, nil
}
