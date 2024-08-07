package main

import(
	route "task05/router"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	route.TaskHandler(router)
	router.Run("localhost:8080")
}