package middleware

import (
	"WePanel/global"
	"WePanel/orm"
	"WePanel/utils/encrypt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func TransactionMiddleware(c *gin.Context) {
	// 开启事务
	tx := global.DB.Begin()

	// 设置事务给请求上下文
	c.Set("tx", tx)

	defer func() {
		// 在请求处理完毕后根据请求结果，提交或回滚事务
		if c.Writer.Status() >= http.StatusInternalServerError {
			// 如果请求发生错误，则回滚事务
			_ = tx.Rollback()
			return
		}

		// 否则提交事务
		_ = tx.Commit()
	}()

	c.Next()
}

type UserDto struct {
	Name      string `json:"name"` //由于返回前端的变量一般都是小写开头，这里规范一下
	Telephone string `json:"telephone"`
}

func ToUserDto(user *orm.User) UserDto {
	return UserDto{
		Name:      user.Username,
		Telephone: user.Telephone,
	}
}

// AuthMiddleware 验证解析token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取authorization header
		tokenString := c.GetHeader("Authorization")

		//验证token格式,若token为空或不是以Bearer开头，则token格式不对
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort() //将此次请求抛弃
			return
		}

		tokenString = tokenString[7:] //token的前面是“bearer”，有效部分从第7位开始

		//从tokenString中解析信息
		token, claims, err := encrypt.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}

		// 查询tokenString中的user信息是否存在
		userId := claims.UserId
		var user orm.User
		global.DB.First(&user, userId)

		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}

		//若存在该用户则将用户信息写入上下文
		userDto := ToUserDto(&user)
		c.Set("user", userDto)
		c.Next()
	}
}
