package main

import (
	_ "fmt"

	"net/http"

	"gihub.com/darkmagic66/a-tour-of-go/control"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	member := control.NewMember()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/getAllData", func(c *gin.Context) {
		c.JSON(http.StatusOK, member.GetAllData())
	})
	router.Run("localhost:3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
