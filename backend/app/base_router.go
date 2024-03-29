package app

import (
	"WePanel/backend/app/middleware"
	"WePanel/backend/app/router"
	"WePanel/backend/global"
	"WePanel/backend/utils/config"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var Router *gin.Engine

func InitRouter(Router *gin.Engine) {
	router.AccountRouter{}.InitRouter(Router)
}

func Init() {
	mode := config.GetConfig("app", "mode")
	port := config.GetConfig("app", "port")
	switch mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	Router = gin.Default()
	Router.Use(middleware.CorsMiddleware())
	Router.Use(middleware.TransactionMiddleware)
	Router.Use(middleware.AuthMiddleware())
	InitRouter(Router)

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      Router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			global.LOG.Fatalf("Serve start failed: %s", err)
		}
	}()
	// 优雅Shutdown（或重启）服务
	// 5秒后优雅Shutdown服务
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt) //syscall.SIGKILL
	<-quit
	global.LOG.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.LOG.Fatalf("Server Shutdown: %s", err)
	}
	select {
	case <-ctx.Done():
	}
	global.LOG.Info("Server exiting")
}
