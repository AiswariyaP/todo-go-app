package routes

import (
	"GOLANG/todo/handler"

	"github.com/gin-gonic/gin"
)

//InitRoutes  defines routes 
func InitRoutes(o *gin.RouterGroup) {
	task := new(handler.TodoController)
	o.POST("/todo", task.CreateTodo)
	o.GET("/todo", task.GetTodo)
	o.PUT("/todo/:id", task.UpdateTodo)
	o.DELETE("/todo/:id", task.DeleteTodo)
	o.GET("/todo/:id", task.GetTodoById)

}
