package store

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

var MgoSession *mgo.Session

const (
	DBTopic         string = "topic"
	DBCollTopic     string = "image_topic"
	DBCollImageData string = "image_data"
)

func SetupMongoDB(url string) {
	c, err := mgo.Dial(url)
	if err != nil {
		fmt.Println("connect mongo fail")
		return
	}
	MgoSession = c
	fmt.Println("connect mongodb success")
}
