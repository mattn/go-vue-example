package config

import (
	"crypto/tls"
	"log"
	"net"
	"net/url"
	"os"

	"github.com/globalsign/mgo"
)

var mongodbSession *mgo.Session
var (
	dbPrefix = "go-vue-example"
)

// DBSession returns the current db session.
func DBSession() *mgo.Session {
	if mongodbSession != nil {
		return mongodbSession
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri = "mongodb://localhost"
	}

	sslmode := false
	hostname := ""
	if url, err := url.Parse(uri); err == nil {
		hostname = url.Hostname()
		values := url.Query()
		if ssl := values.Get("ssl"); ssl != "" {
			if ssl == "true" {
				sslmode = true
			}
			values.Del("ssl")
			url.RawQuery = values.Encode()
			uri = url.String()
		}
	}

	di, err := mgo.ParseURL(uri)
	if err != nil {
		log.Fatalf("Can't parse mongo uri, go error %v\n", err)
	}
	if sslmode {
		di.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.TCPAddr().String(), &tls.Config{
				ServerName: hostname,
			})
			if err != nil {
				log.Println(err)
			}
			return conn, err
		}
	}
	mongodbSession, err = mgo.DialWithInfo(di)
	if mongodbSession == nil || err != nil {
		log.Fatalf("Can't connect to mongo, go error %v\n", err)
	}

	mongodbSession.SetSafe(&mgo.Safe{})
	return mongodbSession
}

// DB returns a database given a name.
func DB(name string) *mgo.Database {
	return DBSession().DB(name)
}

// DefaultDB returns the default database.
func DefaultDB() *mgo.Database {
	switch Environment {
	case "test":
		{
			return DB(dbPrefix + "-test")
		}
	case "production":
		{
			return DB(dbPrefix + "-production")
		}
	}

	return DB(dbPrefix + "-development")
}

// AddBasicIndex add a ascending index given a list of `keys`. The index is always built in background.
func AddBasicIndex(collection *mgo.Collection, keys ...string) {
	collection.EnsureIndex(mgo.Index{
		Key:        keys,
		Background: true,
	})
}

// vi:syntax=go
