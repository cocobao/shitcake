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
		Icon:     "/static/img/test_img_3.png",
		Title:    "测试标题测试标题",
		Msg:      "测试内容测试内容",
		Category: 1,
		IsVip:    false,
		Images:   []string{},
	})
}
