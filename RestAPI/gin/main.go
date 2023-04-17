package main

import (
	_ "fmt"

	"net/http"

	"gihub.com/darkmagic66/a-tour-of-go/control"
	"gihub.com/darkmagic66/a-tour-of-go/model"
	_ "github.com/darkmaic66/a-tour-of-go/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	router := gin.Default()
	member := control.NewMember()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/getAllData", func(c *gin.Context) {
		c.JSON(http.StatusOK, member.GetAllData())
	})
	router.POST("/postCreateData", func(c *gin.Context) {
		var v model.Member
		err := c.ShouldBind(&v)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{
				"message": "error format",
			})
		}
		member.PostCreateData(v)
		c.JSON(http.StatusOK, v)
	})
	router.Run("localhost:3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
