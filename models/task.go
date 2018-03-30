package models

import (
	"errors"
	"net/url"
	"strconv"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/mattn/go-vue-example/helpers"
)

var (
	// ErrInvalidDoc is return when the document is not valid
	ErrInvalidDoc = errors.New("document is invalid")
)

// Task is a model
type Task struct {
	ID        bson.ObjectId  `bson:"_id,omitempty"`
	Body      string         `bson:"body"`
	Done      bool           `bson:"done"`
	CreatedAt int64          `bson:"created_at"`
	UpdatedAt int64          `bson:"updated_at"`
	Errors    helpers.Errors `bson:"-"`
}

// NewTask creates a new instance of Task
func NewTask() *Task {
	task := &Task{}

	// Set default values here

	return task
}

// SaveWithSession saves the Task using the given db session
func (task *Task) SaveWithSession(dbSession *TaskDBSession) error {
	isNewRecord := (len(task.ID) == 0)
	task.setDefaultFields(isNewRecord)

	if !task.IsValid() {
		return ErrInvalidDoc
	}

	if isNewRecord {
		return task.insert(dbSession)
	}

	return task.update(dbSession)
}

// Save saves the Task
func (task *Task) Save() error {
	dbSession := NewTaskDBSession()
	defer dbSession.Close()

	return task.SaveWithSession(dbSession)
}

func (task *Task) DeleteWithSession(dbSession *TaskDBSession) error {
	return dbSession.Delete(bson.M{
		"_id": task.ID,
	})
}

func (task *Task) Delete() error {
	dbSession := NewTaskDBSession()
	defer dbSession.Close()

	return task.DeleteWithSession(dbSession)
}

func (task *Task) setDefaultFields(isNewRecord bool) {
	if isNewRecord {
		task.ID = bson.NewObjectId()
		task.CreatedAt = time.Now().Unix()
	} else {
		task.UpdatedAt = time.Now().Unix()
	}
}

func (task *Task) insert(dbSession *TaskDBSession) error {
	return dbSession.Insert(task)
}

func (task *Task) update(dbSession *TaskDBSession) error {
	return dbSession.Update(task)
}

// ToResponseMap converts the Task to a map to be returned by the API as a response.
func (task *Task) ToResponseMap() helpers.ResponseMap {
	return helpers.ResponseMap{
		"id":   task.ID,
		"body": task.Body,
		"done": task.Done,
	}
}

// ErrorMessages returns validation errors.
func (task *Task) ErrorMessages() helpers.ResponseMap {
	errorMessages := helpers.ResponseMap{}
	for fieldName, message := range task.Errors.Messages {
		errorMessages[fieldName] = message
	}

	return errorMessages
}

// IsValid returns true if all validations are passed.
func (task *Task) IsValid() bool {
	task.Errors.Clear()

	// Run validations here.

	return !task.Errors.HasMessages()
}

// SetAttributes sets the attributes for this Task
func (task *Task) SetAttributes(params url.Values) {
	// for avoiding compile unused strconv error.
	var _ = strconv.FormatBool(true)

	if value, ok := params["body"]; ok {
		task.Body = value[0]
	}
	if value, ok := params["done"]; ok {
		task.Done, _ = strconv.ParseBool(value[0])
	}
}

// FindOneTask finds a task in the database using the given `query`
func FindOneTask(query bson.M) (*Task, error) {
	dbSession := NewTaskDBSession()
	defer dbSession.Close()

	criteria := dbSession.Query(query)
	return FindOneTaskByCriteria(criteria)
}

// FindOneTaskByCriteria finds a single document given a criteria
func FindOneTaskByCriteria(criteria *mgo.Query) (*Task, error) {
	task := &Task{}
	err := criteria.One(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// FindOneTaskByID finds a Task in the database using the given `id`
func FindOneTaskByID(id string) (*Task, error) {
	taskID := bson.ObjectIdHex(id)
	return FindOneTask(bson.M{"_id": taskID})
}

// LoadTasks takes a query a tries to convert it into a Task array.
func LoadTasks(query *mgo.Query) ([]*Task, error) {
	tasks := []*Task{}

	err := query.All(&tasks)
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

// vi:syntax=go
