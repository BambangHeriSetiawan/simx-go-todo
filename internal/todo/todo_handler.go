package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUsecase(c *gin.Context) (TodoUsecase, bool) {
	usecase, ok := c.MustGet("todoUsecase").(TodoUsecase)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "usecase not found"})
	}
	return usecase, ok
}

func GetTodosHandler(c *gin.Context) {
	usecase, ok := getUsecase(c)
	if !ok {
		return
	}
	todos, err := usecase.GetTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func CreateTodoHandler(c *gin.Context) {
	usecase, ok := getUsecase(c)
	if !ok {
		return
	}
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	if err := usecase.CreateTodo(todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "created"})
}

func UpdateTodoHandler(c *gin.Context) {
	usecase, ok := getUsecase(c)
	if !ok {
		return
	}
	id := c.Param("id")
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	if err := usecase.UpdateTodo(id, todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func DeleteTodoHandler(c *gin.Context) {
	usecase, ok := getUsecase(c)
	if !ok {
		return
	}
	id := c.Param("id")
	if err := usecase.DeleteTodo(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
