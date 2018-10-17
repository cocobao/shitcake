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
	return out, coll.Find(nil).Sort("-create_time").Limit(count).Skip(page * count).All(&out)
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

func SetTopicNice(tid string, state int) error {
	c := MgoSession.Clone()
	defer c.Close()
	coll := c.DB(DBTopic).C(DBCollTopic)

	tinfo := &model.ImageTopic{}
	if err := coll.FindId(tid).One(&tinfo); err != nil {
		return err
	}

	var err error
	if state == 0 {
		tinfo.PraiseCount++
		_, err = coll.UpsertId(tid, bson.M{"$set": bson.M{"praise_count": tinfo.PraiseCount}})
	} else if state == 1 {
		tinfo.LowCount++
		_, err = coll.UpsertId(tid, bson.M{"$set": bson.M{"low_count": tinfo.LowCount}})
	}
	return err
}
