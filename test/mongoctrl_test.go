package test

import (
	"testing"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func ConnectMongo() *mgo.Session {
	c, err := mgo.Dial("mongodb://127.0.0.1:27017?maxPoolSize=10")
	if err != nil {
		panic(err)
	}
	return c
}

func TestClearData(t *testing.T) {
	c := ConnectMongo()
	defer c.Close()

	c.DB("topic").C("image_topic").RemoveAll(bson.M{})
}
