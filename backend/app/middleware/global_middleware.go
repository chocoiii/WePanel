package middleware

import (
	"WePanel/backend/app/dto"
	"WePanel/backend/global"
	"WePanel/backend/orm"
	"WePanel/backend/utils/encrypt"
	"WePanel/backend/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CorsMiddleware() gin.HandlerFunc {
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

// AuthMiddleware 验证解析token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取authorization header
		tokenString := c.GetHeader("Authorization")
		//验证token格式,若token为空或不是以Bearer开头，则token格式不对
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			global.LOG.Error("Token's format wrong")
			response.Unauthorized(c)
			c.Abort() //将此次请求抛弃
			return
		}

		tokenString = tokenString[7:] //token的前面是“bearer”，有效部分从第7位开始

		//从tokenString中解析信息
		token, claims, err := encrypt.ParseToken(tokenString)
		if err != nil || !token.Valid {
			global.LOG.Errorf("Parse token failed: %s", err)
			response.Unauthorized(c)
			c.Abort()
			return
		}

		// 查询tokenString中的user信息是否存在
		userId := claims.UserId
		var user orm.User
		global.DB.First(&user, userId)

		if user.ID == 0 {
			global.LOG.Error("User don't existed")
			response.Unauthorized(c)
			c.Abort()
			return
		}
		var userDto = dto.UserDto{
			Username:  user.Username,
			Telephone: user.Telephone,
		}
		//若存在该用户则将用户信息写入上下文
		c.Set("user", userDto)
		c.Next()
	}
}
