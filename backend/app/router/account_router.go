package router

import (
	"WePanel/backend/app/api"
	"github.com/gin-gonic/gin"
)

type AccountRouter struct{}

func (t AccountRouter) InitRouter(Router *gin.Engine) {
	accountRouter := Router.Group("account")
	{
		accountRouter.POST("login", api.AccountApi{}.Login)
		accountRouter.POST("register", api.AccountApi{}.Register)
	}
}
