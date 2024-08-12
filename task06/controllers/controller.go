package controllers

import (
	"net/http"
	data "task06/data"
	models "task06/models"
	"github.com/gin-gonic/gin"
)

func LogIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userInfo models.User
		if err := c.BindJSON(&userInfo); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid format"})
			return
		}
		
		token, err := data.AuthenticateUser(userInfo)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error":"Invalid Credentials"})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{"logged in, here is your token": token})
	}
}

func SignUp() gin.HandlerFunc {
	return func (c *gin.Context) {
		var newUser models.User
		
		if err := c.BindJSON(&newUser); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error":"invalid user format"})
			return
		}
		
		firstUser := data.VerifyFirst()
		if firstUser {
			newUser.Role = "admin"
		} else {newUser.Role = "user"}

		err := data.RegisterUser(newUser)
		
		if err != nil {
			if err.Error() == "neccesary fields are missing" {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
				return
			} else if err.Error() == "user email already in use" {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
				return
			} else {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
				return
			}
		} 

		c.IndentedJSON(http.StatusOK, gin.H{"message":"user registered successfully"})
	}
}

func PromoteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		AuthUser, ok := c.Get("AuthorizedUser")
		if !ok {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error":"Authorization error"})
			return
		}

		AuthorizedUser := AuthUser.(*models.AuthenticatedUser)

		if AuthorizedUser.Role != "admin" {
			c.IndentedJSON(http.StatusForbidden, gin.H{"error":"You are not authorized to promote another user"})
			return
		}

		user_id := c.Param("id")
		statusUpdated := data.UpdateStatus(user_id)
		if !statusUpdated {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error":"user with the specified ID not found"})
			return 
		}

		c.IndentedJSON(http.StatusOK, gin.H{"message": "previlege updated to admin for the specified user"})
	}
}

func AllTasks() gin.HandlerFunc {
	return func (c *gin.Context) {
		allTasks, err := data.FindAllTasks()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error" : "error while fetching data"})
			return
		}
		c.IndentedJSON(http.StatusOK, allTasks)
	}
}

func SpecTask() gin.HandlerFunc {
	return func (c *gin.Context) {
		task_id := c.Param("id")
		foundInfo, err := data.SearchByID(task_id)
		if err!= nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error":"task not found"})
			return
		}

		c.IndentedJSON(http.StatusOK, foundInfo)
	}
}

func UpdateTask() gin.HandlerFunc {

	return func (c *gin.Context) {
		AuthUser, ok := c.Get("AuthorizedUser")
		if !ok {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error":"Authorization error"})
			return
		}

		AuthorizedUser := AuthUser.(*models.AuthenticatedUser)

		if AuthorizedUser.Role != "admin" {
			c.IndentedJSON(http.StatusForbidden, gin.H{"error":"You are not authorized to update a task"})
			return
		}
		task_id := c.Param("id")
		var modifiedTask models.Task 
		if err := c.BindJSON(&modifiedTask); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error":"invalid input format"})
			return
		}

		err := data.ModifyTask(modifiedTask, task_id)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message":"task not found"})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{"message":"Task edited successfully"})

	}
}

func DeleteTask() gin.HandlerFunc {
	return func (c *gin.Context) {
		AuthUser, ok := c.Get("AuthorizedUser")
		if !ok {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error":"Authorization error"})
			return
		}

		AuthorizedUser := AuthUser.(*models.AuthenticatedUser)
		
		if AuthorizedUser.Role != "admin" {
			c.IndentedJSON(http.StatusForbidden, gin.H{"error":"You are not authorized to delete a task"})
			return
		}
	
		task_id := c.Param("id")
		err := data.DeleteByID(task_id)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error":"task not found"})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{"message":"task deleted successfully"})
		
	}
}

func PostTask() gin.HandlerFunc {
	return func (c *gin.Context) {

		AuthUser, ok := c.Get("AuthorizedUser")
		if !ok {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error":"Authorization error"})
			return
		}

		AuthorizedUser := AuthUser.(*models.AuthenticatedUser)

		if AuthorizedUser.Role != "admin" {
			c.IndentedJSON(http.StatusForbidden, gin.H{"error":"You are not authorized to post a task"})
			return
		}
		var newTask models.Task
		if err := c.BindJSON(&newTask); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error":"Invalid request format"})
			return
		}

		err := data.AddTask(newTask)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error":"error while inserting the task to database"})
			return
		}

		c.IndentedJSON(http.StatusOK,gin.H{"message":"Task added successfully"})
	}
}