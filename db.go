package textread

import (
	"log"

	"gopkg.in/mgo.v2"
)

var (
	session *mgo.Session
	url     = "localhost:27017"
	dbname  = "text"
)

func init() {
	var err error
	session, err = mgo.Dial(url)
	if err != nil {
		log.Fatalf("Err connect to mongo: %s", err)
	}
	// db = session.DB(dbname)
}

func SetDB(name string) {
	dbname = name
}

func C(name string) (*mgo.Collection, func()) {
	s := session.Copy()
	c := s.DB(dbname).C(name)
	return c, func() {
		s.Close()
	}
}
