package router

import (
	"github.com/gin-gonic/gin"
	controllers "task04/controllers"
)

func TaskRoutes(internalRouter *gin.Engine) {
	internalRouter.GET("/tasks", controllers.AllTasks())
	internalRouter.GET("/tasks/:id", controllers.SpecTask())
	internalRouter.PUT("/tasks/:id", controllers.UpdateTask())
	internalRouter.DELETE("/tasks/:id", controllers.DeleteTask())
	internalRouter.POST("/tasks", controllers.PostTask())
}
