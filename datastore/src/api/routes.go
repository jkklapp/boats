package api

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

// InitRoutes initializes routes for the backend API
// e.g. example.com/api/v1/healthcheck
func InitRoutes(router *gin.Engine) *gin.Engine {
	// Group of API calls
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/process", process)
		}
	}

	return router
}

func process(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
	// Here the data is processed and saved into a map or database
}
