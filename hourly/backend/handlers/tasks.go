package handlers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/afkjon/Fortress/hourly/backend/db"
	"github.com/afkjon/Fortress/hourly/backend/models"
	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
)

/* API Endpoints */
func handleGetAllTasks(c echo.Context) error {
	var tasks []models.Task
	db.DB.Find(&tasks, "deleted_at is null")
	return c.JSON(http.StatusOK, tasks)
}

func handlePostTasks(c echo.Context) error {
	var newTask models.Task

	if err := c.Bind(&newTask); err != nil {
		return c.HTML(http.StatusNotFound, "No task found!")
	}

	newTask.TargetDate = datatypes.Date(time.Now())

	db.DB.Create(&newTask)

	return c.JSON(http.StatusCreated, newTask)
}

func handleGetTaskById(c echo.Context) error {
	id := c.Param("id")

	task := &models.Task{}
	db.DB.First(&task, id)
	if task != nil {
		return c.JSON(http.StatusOK, task)
	}

	return c.JSON(http.StatusNotFound, "message: Task not found.")
}

func handleDeleteTask(c echo.Context) error {
	id := c.Param("id")

	task := &models.Task{}
	db.DB.First(&task, id)
	if task != nil {
		db.DB.Delete(&task)
		return c.JSON(http.StatusOK, task)
	}

	return c.JSON(http.StatusNotFound, "message: Task not found.")
}

func handleGetTasksByDate(c echo.Context) error {
	date := c.Param("date")

	var tasks []models.Task
	db.DB.Debug().Where("target_date = ?", date).Find(&tasks)

	return c.JSON(http.StatusOK, tasks)
}

func handleExportCsv(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "text/csv")
	c.Response().Header().Set("Content-Disposition", "attachment; filename=output.csv")

	csvWriter := csv.NewWriter(c.Response().Writer)
	defer csvWriter.Flush()

	var tasks []models.Task

	// Get all Tasks and export
	db.DB.Find(&tasks)

	header := tasks[0]
	writeCsvHeader(csvWriter, header)

	for _, t := range tasks {
		row := taskToStrings(t)
		if err := csvWriter.Write(row); err != nil {
			fmt.Printf("Error writing CSV file. %s\n", csvWriter.Error().Error())
			return c.JSON(http.StatusInternalServerError, "Something went wrong.")
		}
	}

	return c.Attachment("output.csv", "output")
}

func writeCsvHeader(w *csv.Writer, data interface{}) {
	var header []string

	typ := reflect.TypeOf(data)
	if typ.Kind() != reflect.Struct {
		fmt.Println("Input is not a struct.")
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		header = append(header, field.Name)
	}

	w.Write(header)
}

func taskToStrings(t models.Task) []string {
	row := []string{strconv.Itoa(int(t.ID)), t.Name, t.Project, strconv.Itoa(int(t.Hours)), t.GetDateString()}
	fmt.Println(row)

	return row
}
