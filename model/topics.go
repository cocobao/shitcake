package model

type ImageTopic struct {
	//主题id
	TopicID string `bson:"topic_id" json:"topic_id"`
	//封面图标
	Icon string `bson:"icon" json:"icon"`
	//标题
	Title string `bson:"title" json:"title"`
	//分类
	Category int `bson:"category" json:"category"`
	//创建时间
	CreateTime string `bson:"create_time" json:"create_time"`
	//是否会员专享
	IsVip string `bson:"is_vip" json:"is_vip"`
	//查看次数
	SeeTime int64 `bson:"see_time" json:"see_time"`
	//主图列表
	Images []string `bson:"images" json:"images"`
}
