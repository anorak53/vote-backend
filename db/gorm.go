package db

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

func GetGormDB() *gorm.DB {
	dbOnce.Do(func() {
		dsn := "root:root@tcp(127.0.0.1:3306)/vote?parseTime=true"
		newDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		db = newDB
		db.AutoMigrate(&Vote{})
		db.AutoMigrate(&User{})
		// VoteC := Vote{
		// 	Name:    "ก้าววันละนิด",
		// 	Details: "พวกเรารักชาติ",
		// 	LogoUrl: "https://www.gstatic.com/webp/gallery/1.jpg",
		// }
		// db.Create(&VoteC)
		// userC := User{
		// 	ID_CARD_NUMBER: 1234,
		// 	STUDENT_NUMBER: 1234,
		// 	IsVoted:        false,
		// }
		// db.Create(&userC)
	})
	return db
}
