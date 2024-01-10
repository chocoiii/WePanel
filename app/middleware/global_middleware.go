package middleware

import (
	"WePanel/global"
	"github.com/gin-gonic/gin"
	"net/http"
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
