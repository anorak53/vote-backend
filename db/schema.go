package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID_CARD_NUMBER int64 `gorm:"unique"`
	STUDENT_NUMBER int64 `gorm:"unique"`
	IsVoted        bool
}

type Vote struct {
	gorm.Model
	Name    string `gorm:"unique"`
	Details string
	LogoUrl string
	Score   int64
}
