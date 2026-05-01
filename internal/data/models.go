package data

import "gorm.io/gorm"

type Models struct {
	User *UserModel
}

func New(db *gorm.DB) *Models {
	return &Models{
		User: &UserModel{
			db: db,
		},
	}
}
