package server

import (
	"WePanel/backend/app"
	"WePanel/backend/utils/db"
	"WePanel/backend/utils/logs"
)

func Start() {
	logs.Init()
	db.Init()
	app.Init()
}
