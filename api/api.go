package api

import (
	router "github.com/gary23w/metasearch_api/api/routers"
)

func Start(port string) {
	router := router.Router()
	router.Run(":" + port)
}
