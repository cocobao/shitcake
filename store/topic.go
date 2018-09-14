package store

import (
	"github.com/cocobao/shitcake/model"
	"gopkg.in/mgo.v2/bson"
)

func InsertTopic(data *model.ImageTopic) error {
	c := MgoSession.Clone()
	defer c.Close()
	coll := c.DB(DBTopic).C(DBColl)

	return coll.Insert(data)
}

func SaveImageTopic(data *model.ImageTopic) error {
	c := MgoSession.Clone()
	defer c.Close()

	selector := bson.M{"topic_id": data.ID}
	coll := c.DB(DBTopic).C(DBColl)
	_, err := coll.Upsert(selector, data)
	return err
}

func GetTopics(count, page int) []*model.ImageTopic {
	c := MgoSession.Clone()
	defer c.Close()
	coll := c.DB(DBTopic).C(DBColl)

	iter := coll.Find(nil).Sort("$ctime:-1").Limit(count).Iter()

	var out []*model.ImageTopic
	var one model.ImageTopic
	for iter.Next(&one) {
		out = append(out, &one)
	}
	return out
}

func GetImageTopicWithTid(topicID string) (*model.ImageTopic, error) {
	c := MgoSession.Clone()
	defer c.Close()
	coll := c.DB(DBTopic).C(DBColl)

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
	coll := c.DB(DBTopic).C(DBColl)

	selector := bson.M{"topic_id": topicID}
	return coll.Remove(selector)
}
