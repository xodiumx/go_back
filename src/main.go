// @title           Example API
// @version         1.0
// @description     Test Swagger for Gin.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.example.com/support
// @contact.email  support@example.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /

package main

import (
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"go_back/api"
	_ "go_back/docs"
	"go_back/logs"
	"go_back/middlewares"
	_ "go_back/settings"
	"log"
)

func main() {

	// Initialize logger
	logs.InitLogger()
	defer logs.CloseLogger()

	// Initialize app
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middlewares.ZapLoggerMiddleware(logs.Logger))

	// Setup router
	base := router.Group("/api")
	{
		v1 := base.Group("/v1")
		{
			// Routes
			v1.GET("/hello", api.HelloHandler)
			v1.GET("/key/:key", api.GetHandler)
			v1.GET("/search", api.SearchHandler)
			v1.POST("/create", api.CreateHandler)
		}
		v2 := base.Group("/v2")
		{
			v2.GET("/hand1", api.Handler1)
			v2.GET("/hand2", api.Handler2)
			v2.GET("/users/:id", api.GetUserHandlerWithSchema)
		}
	}

	// Swagger on
	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	log.Fatal(router.Run(":8080"))
}
