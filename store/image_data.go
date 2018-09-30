package store

import (
	"github.com/cocobao/shitcake/model"
	"gopkg.in/mgo.v2/bson"
)

func InsertImageData(data *model.ImageData) error {
	c := MgoSession.Clone()
	defer c.Close()
	coll := c.DB(DBTopic).C(DBCollImageData)

	return coll.Insert(data)
}

func GetImageData(tid string) ([]*model.ImageData, error) {
	c := MgoSession.Clone()
	defer c.Close()

	coll := c.DB(DBTopic).C(DBCollImageData)

	var data []*model.ImageData
	err := coll.Find(bson.M{"topic_id": tid}).Sort("-insert_time").Limit(10).All(&data)
	return data, err
}
