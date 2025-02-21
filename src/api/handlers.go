package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// HelloHandler @Summary      Получить приветствие
// @Description  Возвращает сообщение "Hello, World!"
// @Tags         example
// @Accept       json
// @Produce      json
// @Success      200  {string}  string  "Hello, World!"
// @Router       /api/v1/hello [get]
func HelloHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello, World!"})
}

// GetHandler - обработчик запроса /v1/get/{key}
// @Summary Получить значение по ключу
// @Description Возвращает переданный ключ
// @Tags example
// @Accept json
// @Produce json
// @Param key path string true "Ключ"
// @Success 200 {object} map[string]string
// @Router /api/v1/key/{key} [get]
func GetHandler(c *gin.Context) {
	fmt.Print("dsadsadsa")
	key := c.Param("key") // Получаем параметр key из URL
	c.JSON(http.StatusOK, gin.H{"key": key})
}

// SearchHandler - обработчик для /v1/search
// @Summary Поиск по query-параметрам
// @Description Возвращает результаты поиска по параметру "q" и отображает выбранную страницу через параметр "page"
// @Tags search
// @Accept  json
// @Produce json
// @Param q query string false "Поисковый запрос"  // Параметр q (необязательный)
// @Param page query int false "Номер страницы"    // Параметр page (необязательный, по умолчанию 1)
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/v1/search [get]
func SearchHandler(c *gin.Context) {
	// Получение query-параметра "q"
	query := c.DefaultQuery("q", "default") // Если параметр не передан, будет использовано "default"
	page := c.DefaultQuery("page", "1")     // Параметр для страницы с дефолтным значением "1"

	// Можно также использовать c.Query() для получения значения параметра без дефолтного значения:
	// q := c.Query("q")  // Вернется пустая строка, если параметр "q" не найден

	c.JSON(http.StatusOK, gin.H{
		"query": query,
		"page":  page,
	})
}

// CreateHandler - обработчик для /v1/create
// @Summary Создать новый объект
// @Description Создает новый объект, принимая JSON в теле запроса
// @Tags create
// @Accept  json
// @Produce json
// @Param body body map[string]interface{} true "Тело запроса"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/v1/create [post]
func CreateHandler(c *gin.Context) {
	var body map[string]interface{} // Структура для тела запроса

	// Связываем тело запроса с переменной
	if err := c.ShouldBindJSON(&body); err != nil {
		// Если произошла ошибка при парсинге JSON
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ответ с принятым JSON
	c.JSON(http.StatusOK, gin.H{
		"received": body,
	})
}

// TEST async
func Handler1(c *gin.Context) {
	time.Sleep(10 * time.Second) // Симуляция длительной операции
	c.JSON(200, gin.H{"message": "handler 1 completed after 10 seconds"})
}
func Handler2(c *gin.Context) {
	time.Sleep(1 * time.Second) // Симуляция длительной операции
	c.JSON(200, gin.H{"message": "handler 2 completed after 1 second"})
}
