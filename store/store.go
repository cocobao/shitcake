package store

import (
	"fmt"

	"github.com/cocobao/shitcake/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type storage struct {
	MgoSession *mgo.Session
}

var Db storage

const (
	DBTopic string = "topic"
	DBColl  string = "image_topic"
)

func SetupMongoDB(url string) {
	c, err := mgo.Dial(url)
	if err != nil {
		fmt.Println("connect mongo fail")
		return
	}
	Db.MgoSession = c
	fmt.Println("connect mongodb success")
}

func (d *storage) SaveImageTopic(data *model.ImageTopic) error {
	c := d.MgoSession.Clone()
	defer c.Close()

	selector := bson.M{"topic_id": data.TopicId}
	coll := c.DB(DBTopic).C(DBColl)
	_, err := coll.Upsert(selector, data)
	return err
}
