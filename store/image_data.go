package store

import "github.com/cocobao/shitcake/model"

func InsertImageData(data *model.ImageData) error {
	c := MgoSession.Clone()
	defer c.Close()
	coll := c.DB(DBTopic).C(DBCollImageData)

	return coll.Insert(data)
}
