package routers

import (
	"github.com/gary23w/metasearch_api/api/controllers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.GET("/health", controllers.Health)
	router.GET("/search/:query", controllers.Search)
	router.GET("/autocomplete/:query", controllers.Autocomplete)
	return router
}
