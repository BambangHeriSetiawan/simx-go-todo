package todo

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, usecase TodoUsecase) {
	group := r.Group("/todos")
	// Middleware to inject usecase into context
	group.Use(func(c *gin.Context) {
		c.Set("todoUsecase", usecase)
		c.Next()
	})
	{
		group.GET("/", GetTodosHandler)
		group.POST("/", CreateTodoHandler)
		group.PUT("/:id", UpdateTodoHandler)
		group.DELETE("/:id", DeleteTodoHandler)
	}
}
