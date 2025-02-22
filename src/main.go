// @title           Example API
// @version         1.0
// @description     Это тестовый сервер Swagger для Gin.
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
	"log"
	"stress/api"
	_ "stress/docs"
)

func main() {
	router := gin.Default()

	base := router.Group("/api")
	{
		v1 := base.Group("/v1") // Вложенная группа для версии API
		{
			// Роуты
			v1.GET("/hello", api.HelloHandler)
			v1.GET("/key/:key", api.GetHandler)
			v1.GET("/search", api.SearchHandler)
			v1.POST("/create", api.CreateHandler)
			v1.GET("/hand1", api.Handler1)
			v1.GET("/hand2", api.Handler2)
		}
	}

	// Подключаем Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	log.Fatal(router.Run(":8080"))
}
