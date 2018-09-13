package model

type InsertTopicReq struct {
	Icon     string   `json:"icon"`
	Title    string   `json:"title"`
	Msg      string   `json:"msg"`
	Category int      `json:"category"`
	IsVip    bool     `json:"is_vip"`
	Images   []string `json:"images"`
}
