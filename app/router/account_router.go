package router

import (
	"WePanel/app/api"
	"github.com/gin-gonic/gin"
)

type AccountRouter struct{}

func (t AccountRouter) InitRouter(Router *gin.Engine) {
	testRouter := Router.Group("account")
	{
		testRouter.GET("login", api.AccountApi{}.Login)
		testRouter.GET("register", api.AccountApi{}.Register)
	}
}
