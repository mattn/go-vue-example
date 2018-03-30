package controllers

import (
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
	"github.com/mattn/go-vue-example/helpers"
	"github.com/mattn/go-vue-example/models"
)

var (
	// ErrTaskNotFound is returned when the task was not found.
	ErrTaskNotFound = helpers.NewResponseError(404, "task not found")

	// ErrCannotDeleteTask is returned when the task can't be deleted.
	ErrCannotDeleteTask = helpers.NewResponseError(404, "cannot delete task")
)

// TasksController is the Task controller.
type TasksController struct {
	Router *echo.Router
}

func (controller *TasksController) createTask(c echo.Context) error {
	c.Request().ParseForm()

	task := &models.Task{}
	task.SetAttributes(c.Request().Form)

	defer task.Save()
	if len(task.ErrorMessages()) == 0 {
		return helpers.JSONResponseObject(c, 200, task)
	}
	return helpers.JSONResponse(c, 400, task.ErrorMessages())
}

func (controller *TasksController) listTasks(c echo.Context) error {
	session := models.NewTaskDBSession()
	defer session.Close()

	query := session.Query(bson.M{})
	query.Sort("-created_at").Limit(20)

	tasks, _ := models.LoadTasks(query)
	tasksResponse := make([]helpers.ResponseMap, len(tasks))
	for i, tasks := range tasks {
		tasksResponse[i] = tasks.ToResponseMap()
	}

	return helpers.JSONResponseArray(c, 200, tasksResponse)
}

func (controller *TasksController) getTask(c echo.Context) error {
	task, err := findTaskByID(c)

	if err != nil {
		return helpers.JSONResponseError(c, err)
	}

	return helpers.JSONResponseObject(c, 200, task)
}

func (controller *TasksController) updateTask(c echo.Context) error {
	task, err := findTaskByID(c)

	if err != nil {
		return helpers.JSONResponseError(c, err)
	}

	task.SetAttributes(c.Request().Form)

	defer task.Save()
	if len(task.ErrorMessages()) == 0 {
		return helpers.JSONResponseObject(c, 200, task)
	}
	return helpers.JSONResponse(c, 400, task.ErrorMessages())
}

func (controller *TasksController) deleteTask(c echo.Context) error {
	task, err := findTaskByID(c)

	if err != nil {
		return helpers.JSONResponseError(c, err)
	}

	if err := task.Delete(); err == nil {
		return helpers.JSONResponse(c, 200, helpers.ResponseMap{})
	}

	return helpers.JSONResponseError(c, ErrCannotDeleteTask)
}

// Setup sets up routes for the Task controller.
func (controller *TasksController) Setup() {
	controller.Router.Add("POST", "/tasks", controller.createTask)
	controller.Router.Add("GET", "/tasks", controller.listTasks)
	controller.Router.Add("GET", "/tasks/:task_id", controller.getTask)
	controller.Router.Add("PUT", "/tasks/:task_id", controller.updateTask)
	controller.Router.Add("DELETE", "/tasks/:task_id", controller.deleteTask)
}

// vi:syntax=go
