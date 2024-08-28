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

	/*
		e.GET("/", func(c echo.Context) error {
			startStr := c.QueryParam("start")
			start, err := strconv.Atoi(startStr)
			if err != nil {
				start = 1
			}

			var tasks []models.Task
			db.DB.Find(&tasks)
			template := "tasks"

			if start == 1 {
				template = "tasks-index"
			}

			return c.Render(http.StatusOK, template, TaskList{
				Start: 1,
				Next:  start + 1,
				More:  true,
				Tasks: tasks,
			})
		})
	*/
}
