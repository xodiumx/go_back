package api

import (
	"github.com/gin-gonic/gin"
)

// @Summary      Получить приветствие
// @Description  Возвращает сообщение "Hello, World!"
// @Tags         example
// @Accept       json
// @Produce      json
// @Success      200  {string}  string  "Hello, World!"
// @Router       /hello [get]
func HelloHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello, World!"})
}
