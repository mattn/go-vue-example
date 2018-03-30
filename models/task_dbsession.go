package models

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/mattn/go-vue-example/config"
)

var (
	// MaxTasksToFetch is the maximum tasks that are fetched from the db.
	MaxTasksToFetch = 50
)

// TaskDBSession handles the connections to the database
type TaskDBSession struct {
	backend *mgo.Session
}

// NewTaskDBSession creates a new database session.
func NewTaskDBSession() *TaskDBSession {
	return &TaskDBSession{backend: config.DBSession().Copy()}
}

// Close closes the database session.
func (dbSession *TaskDBSession) Close() {
	dbSession.backend.Close()
}

// Insert inserts a task in the database.
func (dbSession *TaskDBSession) Insert(task *Task) error {
	return dbSession.collection().Insert(task)
}

// Update updates the db with the given task
func (dbSession *TaskDBSession) Update(task *Task) error {
	return dbSession.collection().UpdateId(task.ID, task)
}

// Query creates a query
func (dbSession *TaskDBSession) Query(query bson.M) *mgo.Query {
	criteria := dbSession.collection().Find(query).Limit(MaxTasksToFetch)

	return criteria
}

// Delete deletes a single task that matches the given query.
func (dbSession *TaskDBSession) Delete(query bson.M) error {
	return dbSession.collection().Remove(query)
}

// DeleteAll deletes all instances of task that match the given query.
func (dbSession *TaskDBSession) DeleteAll(query bson.M) (int, error) {
	info, err := dbSession.collection().RemoveAll(query)
	return info.Removed, err
}

func (dbSession *TaskDBSession) database() *mgo.Database {
	return dbSession.backend.DB(config.DefaultDBName)
}

func (dbSession *TaskDBSession) collection() *mgo.Collection {
	return dbSession.database().C("tasks")
}

// vi:syntax=go
