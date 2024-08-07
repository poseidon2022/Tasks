package controllers

import (
	"net/http"
	data "task05/data"
	task "task05/models"

	"github.com/gin-gonic/gin"
)



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
		task_id := c.Param("id")
		var modifiedTask task.Task 
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
		var newTask task.Task
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