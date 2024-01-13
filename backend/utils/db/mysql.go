package db

import (
	"WePanel/backend/global"
	"WePanel/backend/orm"
	"WePanel/backend/utils/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func AutoMigrate() {
	_ = global.DB.AutoMigrate(&orm.User{})
}

func Init() {
	var username = config.GetConfig("Mysql", "username")
	var passwd = config.GetConfig("Mysql", "passwd")
	var ip = config.GetConfig("Mysql", "ip")
	var port = config.GetConfig("Mysql", "port")
	var database = config.GetConfig("Mysql", "database")
	dsn := username + ":" + passwd + "@tcp(" + ip + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		global.LOG.Fatalf("mysql connect failed: %s", err)
	}
	global.DB = db
	AutoMigrate()
}
