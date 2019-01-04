package main

import (
	"github.com/lupengyu/aisgin/dal/mysql"
	"github.com/lupengyu/aisgin/dal/redis"
	"github.com/lupengyu/aisgin/handler"
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func initLog() {
	gin.DisableConsoleColor()
	f, _ := os.Create("/log/" + string(time.Now().Unix()) + ".log")
	gin.DefaultWriter = io.MultiWriter(f)
}

func initRouter() *gin.Engine {
	r := gin.Default()

	Api := r.Group("/api/")
	{
		ApiGet := Api.Group("/get/")
		{
			ApiGet.GET("/ship", handler.GetShipID)
			ApiGet.GET("/position/:id", handler.GetPosition)
			ApiGet.GET("/info/:id", handler.GetShipInfo)
		}
	}

	return r
}

func main() {
	defer func() {
		log.Println("Service stopped.")
	}()
	log.Println("Service starting ......")
	initLog()
	mysql.InitMysql()
	redis.InitRedis()
	r := initRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
