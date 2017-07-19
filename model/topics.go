package model

type ImageTopic struct {
	TopicId    string `bson:"topic_id" json:"-"`
	Title      string `bson:"title" json:"title"`
	TopicType  int    `bson:"topic_type" json:"topic_type"`
	CreateTime string `bson:"create_time" json:"create_time"`
	ImageName  string `bson:"image_name" json:"image_name"`
	PathName   string `bson:"path_name" json:"-"`
}
