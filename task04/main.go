package main

import (
	"github.com/gin-gonic/gin"
	routers "task04/router"
)

func main() {

	router := gin.Default()
	routers.TaskRoutes(router)
	router.Run("localhost:8080")
}