package controller

import (
	"net/http"

	log "github.com/cihub/seelog"
	"github.com/cocobao/shitcake/store"
	"github.com/cocobao/shitcake/utils"
	"github.com/gin-gonic/gin"
)

type TopicController struct {
	BaseController
}

func NewTopicController(c *gin.Context) *TopicController {
	ctrl := &TopicController{}
	ctrl.ginCtx = c
	return ctrl
}

func (c *TopicController) Get() {
	tid := c.ginCtx.Query("tid")

	topic, err := store.Db.GetImageTopicWithTid(tid)
	if err != nil {
		c.ginCtx.Redirect(301, "/")
		return
	}
	log.Debugf("topic:%#v", topic)

	c.ginCtx.HTML(http.StatusOK, "pic_detail.html", gin.H{
		"files": utils.StructToMapJson(topic),
	})
}
