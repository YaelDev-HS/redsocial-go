package data

import (
	"errors"
	"strings"
	"time"

	"github.com/YaelDev-HS/redsocial-go/internal/validator"
	"golang.org/x/crypto/bcrypt"
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

func (m *UserModel) FindByEmail(email string) (*User, error) {
	var user User

	err := m.db.Where("email = ? AND is_enabled = TRUE", email).First(&user).Error

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, ErrModelNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (p *User) SetPassword(plainText string) error {
	password, err := bcrypt.GenerateFromPassword([]byte(plainText), 8) // 12

	if err != nil {
		return err
	}

	p.Password = password
	return nil
}

func (u *User) ComparePassword(plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(u.Password, []byte(plainText))

	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func CheckPasswordAndEmail(v *validator.Validator, password, email string) {
	v.Check(len(password) > 3, "password", "is too short")
	v.Check(len(password) < 60, "password", "is too long")
	v.Check(!v.Match(email, validator.EmailRegex), "email", "is not valid")
}
