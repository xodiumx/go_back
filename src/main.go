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
	"stress/src/api"
	_ "stress/src/docs"
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
		}
	}

	// Подключаем Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	router.Run(":8080") // Запускаем сервер
}
