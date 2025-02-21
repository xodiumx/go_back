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
	r := gin.Default()

	// Роуты
	r.GET("/hello", api.HelloHandler)

	// Подключаем Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	r.Run(":8080") // Запускаем сервер
}
