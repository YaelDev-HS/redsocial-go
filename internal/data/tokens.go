package data

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ScopeAuthentication = "AUTHENTICATION"
)

type Token struct {
	ID        int64     `json:"-"`
	UserID    int64     `json:"-"`
	Hash      []byte    `json:"-" gorm:"column:token"`
	Scope     string    `json:"-"`
	PlainText string    `json:"token" gorm:"-"`
	ExpiryAt  time.Time `json:"expiry_at"`
	User      *User     `json:"-" gorm:"foreignKey:UserID"`
}

func (*Token) TableName() string {
	return "tokens"
}

type TokenModel struct {
	db *gorm.DB
}

func generateToken(userID int64, ttl time.Duration, scope string) (*Token, error) {
	token := &Token{
		UserID:   userID,
		ExpiryAt: time.Now().Add(ttl),
		Scope:    scope,
	}

	randomBytes := make([]byte, 16) // [, , , , , , , , , , ,] = [0101110, 1010111]

	// sistema operativo
	_, err := rand.Read(randomBytes)

	if err != nil {
		return nil, err
	}

	//1. Base32.StdEncoding = texto legible, lo pasa A-Z y numeros legibles por humanos = AHAWDWHAJDAGWDHADA
	//2. withpadding = evita signos = y / para que solo tenga letras y numeros
	token.PlainText = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)

	// crea un arreglo de 32 bytes para evitar que sea legible nuevamente
	hash := sha256.Sum256([]byte(token.PlainText))

	// SIEMPRE VIENEN de 26 caracteres
	token.Hash = hash[:] // pasarlo de nuevo a un arreglo de bytes

	return token, err
}

func (m *TokenModel) Insert(userID int64, ttl time.Duration, scope string) (*Token, error) {
	token, err := generateToken(userID, ttl, scope)

	if err != nil {
		return nil, err
	}

	err = m.db.Create(token).Error

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (m *TokenModel) FindByPlaintext(plaintext, scope string) (*Token, error) {
	var token Token
	hash := sha256.Sum256([]byte(plaintext))

	err := m.db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "email", "nickname")
	}).Where("token = ? AND expiry_at > ? AND scope = ?", hash[:], time.Now(), scope).First(&token).Error

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, ErrModelNotFound
		default:
			return nil, err
		}
	}

	return &token, nil
}
