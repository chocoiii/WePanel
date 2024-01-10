package server

import (
	"WePanel/app"
	"WePanel/utils/db"
	"WePanel/utils/logs"
)

func Start() {
	logs.Init()
	db.Init()
	app.Init()
}
