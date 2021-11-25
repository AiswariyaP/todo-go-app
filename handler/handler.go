package handler

import (
	"GOLANG/todo/forms"
	"GOLANG/todo/todoDAO"

	"github.com/gin-gonic/gin"
)

type TodoController struct{}

var taskModel = new(todoDAO.TaskModel)

// CreateTodo creates new todo
func (task *TodoController) CreateTodo(c *gin.Context) {
	var data forms.CreateTaskCommand
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}
	err := taskModel.Create(data)
	if err != nil {
		c.JSON(406, gin.H{"message": "Task could not be created", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "Task Created", "form": data})
}

// GetTodo fetches all todo
func (task *TodoController) GetTodo(c *gin.Context) {

	tasks, err := taskModel.Get()
	if err != nil {
		c.JSON(406, gin.H{"message": "Task could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Tasks Fetched", "form": tasks})
}

// UpdateTodo UPDATE todo by id
func (task *TodoController) UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	data := forms.CreateTaskCommand{}

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Invalid Parameters"})
		c.Abort()
		return
	}

	err := taskModel.Update(id, data)
	if err != nil {
		c.JSON(406, gin.H{"message": "Task Could Not Be Updated", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "Task Updated"})
}

// GetTodoById FETCH todo by id
func (task *TodoController) GetTodoById(c *gin.Context) {
	id := c.Param("id")

	tasks, err := taskModel.GetTodoByIdDAO(id)
	if err != nil {
		c.JSON(406, gin.H{"message": "Task Could Not Be Deleted", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "Task Deleted", "form": tasks})
}

// GetTodoById DELETE todo by id
func (task *TodoController) DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	err := taskModel.Delete(id)
	if err != nil {
		c.JSON(406, gin.H{"message": "Task Could Not Be Deleted", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "Task Deleted"})
}
