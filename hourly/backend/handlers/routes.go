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
	e.GET("/data/tasks", handleGetAllTasks)
	e.POST("/data/tasks", handlePostTasks)
	e.GET("/data/tasks/:id", handleGetTaskById)
	e.GET("/data/tasks/csv", handleExportCsv)
	e.GET("/data/tasks/date/:date", handleGetTasksByDate)
	e.POST("/data/tasks/delete/:id", handleDeleteTask)

	e.POST("/login", Login)
	e.POST("/register", Register)
	e.POST("/logout", Logout)
}
