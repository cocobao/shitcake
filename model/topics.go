package model

type ImageTopic struct {
	TopicID    int64  `bson:"topic_id" json:"topic_id"`
	Title      string `bson:"title" json:"title"`
	Category   int    `bson:"category" json:"category"`
	CreateTime string `bson:"create_time" json:"create_time"`
	IsVip      string `bson:"is_vip" json:"is_vip"`
	SeeTime    int64  `bson:"see_time"`
}
