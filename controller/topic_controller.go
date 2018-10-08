package controller

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/cocobao/log"
	"github.com/cocobao/shitcake/store"
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
	log.Debug("get topic:", tid)

	topic, err := store.GetImageTopicWithTid(tid)
	if err != nil {
		log.Warnf("get topic:%s fail, err:%v", tid, err)
		c.ginCtx.Redirect(301, "/")
		return
	}
	log.Debugf("topic detail:%#v", topic)

	c.ginCtx.HTML(http.StatusOK, "pic_detail.html", gin.H{
		"files": topic,
	})
}

func (c *TopicController) Del() {
	tid := c.ginCtx.Query("tid")
	err := store.DelImageTopicWithTid(tid)
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

func GetTopicDetail(g *gin.Context) {
	tid := g.Param("tid")
	page, _ := strconv.Atoi(g.Param("page"))

	log.Debugf("get topic:%s, page:%s", tid, page)

	imgData, err := store.GetImageData(tid)
	if err != nil {
		log.Warnf("get topic:%s fail, err:%v", tid, err)
		g.Redirect(301, "/")
		return
	}

	imgData, err := store.GetImageData(tid, page, 10)
	if err != nil {
		log.Errorf("get image data fail,%v", err)
		g.Redirect(301, "/")
		return
	}

	log.Debugf("imgData:%+v", imgData)

	g.HTML(http.StatusOK, "pic_detail.html", gin.H{
		"items": topic,
		"imgs":  imgData,
	})
}
