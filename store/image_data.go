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

func GetImageData(tid string, page, limit int) ([]*model.ImageData, error) {
	c := MgoSession.Clone()
	defer c.Close()

	coll := c.DB(DBTopic).C(DBCollImageData)

	var imgData []*model.ImageData
	err := coll.Find(bson.M{"topic_id": tid}).Sort("index").Skip(page * limit).Limit(limit).All(&imgData)
	return imgData, err
}
