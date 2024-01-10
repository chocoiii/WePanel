package api

import (
	"WePanel/orm"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TestApi struct{}

func (t TestApi) TestFunc(c *gin.Context) {
	var db *gorm.DB
	if tx, exist := c.Get("tx"); exist {
		tx, _ := tx.(*gorm.DB)
		db = tx
	}
	user := orm.User{
		Name: "111",
	}
	err := db.Create(&user).Error
	var message string
	if err != nil {
		message = err.Error()
	} else {
		message = "Hello World!"
	}
	c.JSON(200, gin.H{
		"message": message,
	})
}
