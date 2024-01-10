package db

import (
	"WePanel/global"
	"WePanel/utils/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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
}

//func Add(db *gorm.DB, model interface{}, createMap map[string]interface{}) error{
//	err := db.Model(model).Create(createMap).Error
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func BatchAdd(db *gorm.DB, model interface{}, createMap []map[string]interface{}) error{
//	err := db.Model(model).Create(createMap).Error
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func Update(db *gorm.DB, model interface{}, conditions map[string]interface{}, updates map[string]interface{}) error{
//	conditionsExpressions := expressions.NewBuilder().Assignments(conditions)
//	err := db.Model(model).Where(conditionsExpressions).Updates(updates)
//	if err != nil {
//		return err
//	}
//	return nil
//}
