package controllers

import (
	"github.com/labstack/echo"
	"github.com/mattn/go-vue-example/helpers"
	"github.com/mattn/go-vue-example/models"
)

func findTaskByID(c echo.Context) (*models.Task, *helpers.ResponseError) {
	c.Request().ParseForm()

	task, _ := models.FindOneTaskByID(c.Param("task_id"))
	if task == nil {
		return nil, ErrTaskNotFound
	}

	return task, nil
}

// vi:syntax=go
