package helpers

import (
	"github.com/globalsign/mgo"
	"github.com/sirupsen/logrus"
)

// HandleDBError handles the error and reconnects if needed.
func HandleDBError(err error) {
	logrus.WithFields(logrus.Fields{"err": err}).Error("DB Error")
}

// AddBasicIndex add a ascending index given a list of `keys`. The index is always built in background.
func AddBasicIndex(collection *mgo.Collection, keys ...string) {
	err := collection.EnsureIndex(mgo.Index{
		Key:        keys,
		Background: true,
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err, "keys": keys}).Error("Error adding index.")
	}
}

// AddUniqueBasicIndex adds a unique indexes to the given `keys`
func AddUniqueBasicIndex(collection *mgo.Collection, keys ...string) {
	err := collection.EnsureIndex(mgo.Index{
		Key:        keys,
		Background: true,
		Unique:     true,
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err, "keys": keys}).Error("Error adding unique index.")
	}
}

// vi:syntax=go
