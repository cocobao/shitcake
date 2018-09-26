package store

import (
	"github.com/cocobao/shitcake/model"
	"gopkg.in/mgo.v2/bson"
)

func InsertTopic(data *model.ImageTopic) error {
	c := MgoSession.Clone()
	defer c.Close()
	coll := c.DB(DBTopic).C(DBCollTopic)

	return coll.Insert(data)
}

func SaveImageTopic(data *model.ImageTopic) error {
	c := MgoSession.Clone()
	defer c.Close()

	selector := bson.M{"topic_id": data.ID}
	coll := c.DB(DBTopic).C(DBCollTopic)
	_, err := coll.Upsert(selector, data)
	return err
}

func GetTopics(count, page int) ([]*model.ImageTopic, error) {
	c := MgoSession.Clone()
	defer c.Close()
	coll := c.DB(DBTopic).C(DBCollTopic)

	var out []*model.ImageTopic
	return out, coll.Find(nil).Sort("$ctime:-1").Limit(count).Skip(page).All(&out)
}

func GetImageTopicWithTid(topicID string) (*model.ImageTopic, error) {
	c := MgoSession.Clone()
	defer c.Close()
	coll := c.DB(DBTopic).C(DBCollTopic)

	selector := bson.M{"_id": topicID}
	var data model.ImageTopic
	err := coll.Find(selector).One(&data)
	if err != nil {
		return nil, err
	}

	data.SeeCount++
	coll.Upsert(selector, data)
	return &data, nil
}

func DelImageTopicWithTid(topicID string) error {
	c := MgoSession.Clone()
	defer c.Close()
	coll := c.DB(DBTopic).C(DBCollTopic)

	selector := bson.M{"topic_id": topicID}
	return coll.Remove(selector)
}
