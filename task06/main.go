package main

import (
	"github.com/gin-gonic/gin"
	route "task06/router"
)

func main() {

	router := gin.Default()
	route.TaskHandler(router)
	router.Run("localhost:8080")

}