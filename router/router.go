package router

import "vote.app/m/db"

func init() {
	_ = db.GetGormDB()
}
