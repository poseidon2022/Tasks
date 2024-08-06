package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	task "task04/models"
	data "task04/data"
)



func AllTasks() gin.HandlerFunc {
	return func (c *gin.Context) {
		c.IndentedJSON(http.StatusOK, task.Pre)
	}
}

func SpecTask() gin.HandlerFunc {
	return func (c *gin.Context) {
		task_id := c.Param("id")
		foundInfo, err := data.SearchByID(task_id)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message":"task not found"})
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
		id := c.Param("id")
		err := data.DeleteByID(id)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "task not found"})
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
			c.IndentedJSON(http.StatusNotAcceptable, gin.H{"error":"A different task with the same ID exists"})
			return
		}

		c.IndentedJSON(http.StatusOK,gin.H{"message":"Task added successfully"})
	}
}