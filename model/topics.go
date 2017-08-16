package model

type ImageTopic struct {
	TopicID    string   `bson:"topic_id" json:"topic_id"`
	Icon       string   `bson:"icon" json:"icon"`
	Title      string   `bson:"title" json:"title"`
	Category   int      `bson:"category" json:"category"`
	CreateTime string   `bson:"create_time" json:"create_time"`
	IsVip      string   `bson:"is_vip" json:"is_vip"`
	SeeTime    int64    `bson:"see_time" json:"see_time"`
	Images     []string `bson:"images" json:"images"`
}
