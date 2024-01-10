package router

import (
	"WePanel/app/api"
	"github.com/gin-gonic/gin"
)

type TestRouter struct{}

func (t TestRouter) InitRouter(Router *gin.Engine) {
	testRouter := Router.Group("test")
	{
		testRouter.GET("init", api.TestApi{}.TestFunc)
	}
}
