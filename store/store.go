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

	selector := bson.M{"topic_id": data.TopicID}
	coll := c.DB(DBTopic).C(DBColl)
	_, err := coll.Upsert(selector, data)
	return err
}

func (d *storage) GetImageTopic(count int) []model.ImageTopic {
	c := d.MgoSession.Clone()
	defer c.Close()
	coll := c.DB(DBTopic).C(DBColl)

	iter := coll.Find(nil).Sort("$ctime:-1").Limit(count).Iter()

	var out []model.ImageTopic
	var one model.ImageTopic
	for iter.Next(&one) {
		out = append(out, one)
	}
	return out
}

func (d *storage) GetImageTopicWithTid(topicID string) (*model.ImageTopic, error) {
	c := d.MgoSession.Clone()
	defer c.Close()
	coll := c.DB(DBTopic).C(DBColl)

	selector := bson.M{"topic_id": topicID}
	var data model.ImageTopic
	err := coll.Find(selector).One(&data)
	if err != nil {
		return nil, err
	}

	data.SeeTime += 1
	coll.Upsert(selector, data)
	return &data, nil
}

func (d *storage) DelImageTopicWithTid(topicID string) error {
	c := d.MgoSession.Clone()
	defer c.Close()
	coll := c.DB(DBTopic).C(DBColl)

	selector := bson.M{"topic_id": topicID}
	return coll.Remove(selector)
}
