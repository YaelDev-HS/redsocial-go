package data

import "gorm.io/gorm"

type Models struct {
	User    *UserModel
	Token   *TokenModel
	Message *MessageModel
}

func New(db *gorm.DB) *Models {
	return &Models{
		User: &UserModel{
			db: db,
		},
		Token: &TokenModel{
			db: db,
		},
		Message: &MessageModel{
			db: db,
		},
	}
}
