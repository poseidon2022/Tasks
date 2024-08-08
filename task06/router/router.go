package router

import (
	"github.com/gin-gonic/gin"
	controllers "task06/controllers"
	middleware "task06/middleware"
)

func TaskHandler (internalRouter *gin.Engine) {
	internalRouter.GET("/tasks", middleware.UserAuth() ,controllers.AllTasks())
	internalRouter.GET("/tasks/:id", middleware.UserAuth(), controllers.SpecTask())
	internalRouter.PUT("/tasks/:id", controllers.UpdateTask())
	internalRouter.DELETE("/tasks/:id", controllers.DeleteTask())
	internalRouter.POST("/tasks", controllers.PostTask())
	internalRouter.POST("/register", controllers.SignUp())
	internalRouter.POST("/login", controllers.LogIn())
}