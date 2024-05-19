package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	STUDENT_NUMBER int64 `gorm:"unique"`
	IsVoted        bool
}

type Vote struct {
	gorm.Model
	Name    string `gorm:"unique"`
	Number  int64  `gorm:"unique"`
	Details string
	LogoUrl string
	Score   int64
}
