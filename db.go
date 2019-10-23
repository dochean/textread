package textread

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	session *mgo.Session
	url     = "localhost:27017"
	dbname  = "text"
	cname   = "nodes"
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

func Insert(n *Novel) error {
	c, cl := C(cname)
	defer cl()
	if count, err := c.Find(bson.M{
		"title": n.Title,
	}).Count(); count > 0 || err != nil {
		return fmt.Errorf("Err existed title or [%w]", err)
	}
	return c.Insert(n)
}

type Novellist struct {
	Id    bson.ObjectId `bson:"_id"`
	Title string        `bson:"title"`
}

func List() []*Novellist {
	c, cl := C(cname)
	defer cl()
	var list []*Novellist
	err := c.Find(nil).Select(bson.M{"title": 1, "_id": 1}).All(&list)
	if err != nil {
		log.Printf("Err when get novel list from mongo [%s]\n", err)
		return nil
	}
	return list
}

func GetById(id bson.ObjectId) *Novel {
	c, cl := C(cname)
	defer cl()
	var novel *Novel
	err := c.Find(bson.M{"_id": id}).One(&novel)
	if err != nil {
		log.Printf("Err when get novel [id-%s][%s]\n", id, err)
		return nil
	}
	return novel
}
