package controller

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/cocobao/log"
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

func (c *TopicController) Del() {
	tid := c.ginCtx.Query("tid")
	err := store.Db.DelImageTopicWithTid(tid)
	if err != nil {
		log.Warn("del db topic fail,", err)
		return
	}
	pt := path.Join("static/icon", fmt.Sprintf("/%s/", tid))
	os.RemoveAll(pt)
	pt = path.Join("static/images", fmt.Sprintf("/%s/", tid))
	os.RemoveAll(pt)
	c.ginCtx.Redirect(301, "/")
}
