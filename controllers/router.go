package controllers

import (
	"github.com/labstack/echo"
)

// Setup sets up all controllers.
func Setup(router *echo.Router) {
	tasks := TasksController{Router: router}
	tasks.Setup()
}

// vi:syntax=go
