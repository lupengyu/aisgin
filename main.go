package main

import (
	"github.com/lupengyu/aisgin/dal/mysql"
	"github.com/lupengyu/aisgin/handler"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

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
		}
	}
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	defer func() {
		log.Println("Service stopped.")
	}()
	log.Println("Service starting ......")
	initLog()
	mysql.InitMysql()
	r := initRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
