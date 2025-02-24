package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// HelloHandler @Summary
// @Description  Returns "Hello, World!"
// @Tags         example
// @Accept       json
// @Produce      json
// @Success      200  {string}  string  "Hello, World!"
// @Router       /api/v1/hello [get]
func HelloHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello, World!"})
}

// GetHandler - /v1/get/{key}
// @Summary Get key from path param
// @Description Return key from path param
// @Tags example
// @Accept json
// @Produce json
// @Param key path string true "Key"
// @Success 200 {object} map[string]string
// @Router /api/v1/key/{key} [get]
func GetHandler(c *gin.Context) {
	fmt.Print("dsadsadsa")
	key := c.Param("key")
	c.JSON(http.StatusOK, gin.H{"key": key})
}

// SearchHandler - /v1/search
// @Summary Get query params
// @Description Get "q" and "page" query params
// @Tags search
// @Accept  json
// @Produce json
// @Param q query string false "Search query"  // Param q (not required, by default - base)
// @Param page query int false "Pagination"    // Param page (not required, by default - 1)
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/v1/search [get]
func SearchHandler(c *gin.Context) {
	query := c.DefaultQuery("q", "base")
	page := c.DefaultQuery("page", "1")

	// Can use without default:
	// q := c.Query("q")  // If "q" not find default value blank string ""

	c.JSON(http.StatusOK, gin.H{
		"query": query,
		"page":  page,
	})
}

// CreateHandler -/v1/create
// @Summary Data from body
// @Description Received json body
// @Tags create
// @Accept  json
// @Produce json
// @Param body body map[string]interface{} true "Body"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/v1/create [post]
func CreateHandler(c *gin.Context) {
	var body map[string]interface{} // Blank struct of body, can receive any body

	// Deserialize body
	if err := c.ShouldBindJSON(&body); err != nil {
		// Error if not correct json body
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"received": body,
	})
}

// TEST async
func Handler1(c *gin.Context) {
	time.Sleep(10 * time.Second) // Simulation of work
	c.JSON(200, gin.H{"message": "handler 1 completed after 10 seconds"})
}
func Handler2(c *gin.Context) {
	time.Sleep(1 * time.Second) // Simulation of work
	c.JSON(200, gin.H{"message": "handler 2 completed after 1 second"})
}

// GetUserHandlerWithSchema - /api/v2/users/{id}
// @Summary Endpoint with schema
// @Description description
// @Tags with_schema
// @Accept  json
// @Produce json
// @Param id path int true "User id"
// @Success 200 {object} UserResponse
// @Failure 400 {object} BadRequest
// @Router /api/v2/users/{id} [get]
func GetUserHandlerWithSchema(c *gin.Context) {
	user := UserResponse{
		ID:    1,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	c.JSON(http.StatusOK, user)
}
