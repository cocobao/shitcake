package model

type ImageTopic struct {
	ID          string   `bson:"_id" json:"id"`                    //id
	Icon        string   `bson:"icon" json:"icon"`                 //封面图标
	Title       string   `bson:"title" json:"title"`               //标题
	Category    int      `bson:"category" json:"category"`         //分类
	CreateTime  string   `bson:"create_time" json:"create_time"`   //创建时间
	IsVip       bool     `bson:"is_vip" json:"is_vip"`             //是否会员专享
	SeeCount    int64    `bson:"see_count" json:"see_count"`       //查看次数
	PraiseCount int64    `bson:"praise_count" json:"praise_count"` //被赞次数
	Images      []string `bson:"images" json:"images"`             //主图列表
	Description string   `bson:"description" json:"description"`   //说明
}
