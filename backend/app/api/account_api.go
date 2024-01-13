package api

import (
	"WePanel/backend/global"
	"WePanel/backend/orm"
	"WePanel/backend/utils/encrypt"
	"WePanel/backend/utils/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AccountApi struct{}

func (t AccountApi) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	var db *gorm.DB
	if tx, exist := c.Get("tx"); exist {
		tx, _ := tx.(*gorm.DB)
		db = tx
	}
	var user orm.User
	db.Where("username=?", username).First(&user)
	if user.ID == 0 {
		global.LOG.Errorf("[%s] login failed: username not exist", username)
		response.Fail(c, nil, "用户名不存在")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		global.LOG.Errorf("[%s] login failed: password wrong: %s", username, err)
		response.Fail(c, nil, "密码错误")
		return
	}
	token, err := encrypt.GetToken(user)
	if err != nil {
		global.LOG.Errorf("[%s] login failed: generate token failed: %s", username, err)
		response.Fail(c, nil, "token获取异常")
		return
	}
	global.LOG.Infof("[%s] login successfully", username)
	response.Success(c, gin.H{"token": token}, "登录成功")
}

func (t AccountApi) Register(c *gin.Context) {
	username := c.PostForm("username")
	nickname := c.PostForm("nickname")
	password := c.PostForm("password")
	telephone := c.PostForm("telephone")
	var db *gorm.DB
	if tx, exist := c.Get("tx"); exist {
		tx, _ := tx.(*gorm.DB)
		db = tx
	}
	if len(telephone) != 11 {
		global.LOG.Errorf("[%s] register failed: telephone must be 11", username)
		response.Fail(c, nil, "手机号必须11位")
		return
	}
	var user orm.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		global.LOG.Errorf("[%s] register failed: telephone has been registered", username)
		response.Fail(c, nil, "手机号已被注册")
		return
	}
	db.Where("username=?", username).First(&user)
	if user.ID != 0 {
		global.LOG.Errorf("[%s] register failed: username has been registered", username)
		response.Fail(c, nil, "用户名已被注册")
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		global.LOG.Errorf("[%s] register failed: entrypt failed: %s", username, err)
		response.Fail(c, nil, "加密错误")
		return
	}
	addUser := orm.User{
		Username:  username,
		Nickname:  nickname,
		Password:  string(hashedPassword),
		Telephone: telephone,
	}
	if err := db.Create(&addUser).Error; err != nil {
		global.LOG.Errorf("[%s] register failed: account create failed: %s", username, err)
		response.Fail(c, nil, err.Error())
		return
	}
	global.LOG.Infof("[%s] register successfully", username)
	response.Success(c, nil, "注册成功")
}
