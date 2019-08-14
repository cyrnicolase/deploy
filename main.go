package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"deploy/config"
	"deploy/models"
	"deploy/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := bootstrap()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); nil != err && http.ErrServerClosed != err {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shotdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 0*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); nil != err {
		log.Fatal("server shutdown: ", err)
	}

	log.Println("server exiting")
}

// 启动
func bootstrap() (router *gin.Engine) {
	initConfigPath()

	config.Boot()
	models.Boot()
	if "production" == config.App.RunMode {
		gin.SetMode(gin.ReleaseMode)
	}
	router = routes.Boot()

	return
}

// 初始化配置文件路径
func initConfigPath() {
	currentDir, _ := os.Getwd()
	path := currentDir + "/config/yaml"
	config.SetConfigPath(path)
}
