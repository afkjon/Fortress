package handlers

import (
	"github.com/afkjon/Fortress/hourly/backend/models"
	"github.com/labstack/echo/v4"
)

type TaskList struct {
	Start int
	Next  int
	More  bool
	Tasks []models.Task
}

func SetupRoutes(e *echo.Echo) {
	/* API */
	e.GET("/data/tasks", getAllTasks)
	e.POST("/data/tasks", createTask)
	e.GET("/data/tasks/:id", getTaskById)
	e.GET("/data/tasks/csv", exportCsv)
	e.GET("/data/tasks/date/:date", getTasksByDate)
	e.POST("/data/tasks/delete/:id", deleteTask)

	e.POST("/login", Login)
	e.POST("/register", Register)
	e.POST("/logout", Logout)
}
