package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rina-makita0320/moving-checklist/internal/database"
	"github.com/rina-makita0320/moving-checklist/internal/handler"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	tasks := r.Group("/tasks")
	{
		tasks.GET("", handler.GetTasks)
		tasks.POST("", handler.CreateTask)
		tasks.PATCH("/:id/complete", handler.CompleteTask)
	}

	r.Run(":8080")
}