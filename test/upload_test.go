package test

import (
	"testing"

	"github.com/cocobao/shitcake/model"
	"github.com/cocobao/shitcake/utils"
)

var (
	URL_H string = "http://127.0.0.1:8080/"
)

func TestInsertTopic(t *testing.T) {
	utils.DoHttpPostSimple(URL_H+"insert-topic", &model.InsertTopicReq{
		Icon:     "/static/img/capImgs/1/0/1_09292043529D5.jpg",
		Title:    "测试标题测试标题",
		Msg:      "测试内容测试内容",
		Category: 1,
		IsVip:    false,
		Images: []string{
			"/static/img/capImg/1/0/204352E00-34.jpg",
			"/static/img/capImg/1/0/204352IG-0.jpg",
			"/static/img/capImg/1/0/204352MJ-35.jpg",
			"/static/img/capImg/1/0/2043521F6-13.jpg",
		},
	})
}
