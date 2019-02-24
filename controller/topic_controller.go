package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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
		c.ginCtx.Redirect(302, "/")
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
	// pt := path.Join("static/icon", fmt.Sprintf("/%s/", tid))
	// os.RemoveAll(pt)
	// pt = path.Join("static/images", fmt.Sprintf("/%s/", tid))
	// os.RemoveAll(pt)
	c.ginCtx.Redirect(302, "/")
}

func GetTopicDetail(g *gin.Context) {
	tid := g.Param("tid")
	page, _ := strconv.Atoi(g.Param("page"))

	log.Debugf("get topic:%s, page:%s", tid, page)

	topic, err := store.GetImageTopicWithTid(tid)
	if err != nil {
		log.Warnf("get topic:%s fail, err:%v", tid, err)
		g.Redirect(302, "/")
		return
	}
	num := 10
	imgData, err := store.GetImageData(tid, page, num)
	if err != nil {
		log.Errorf("get image data fail,%v", err)
		g.Redirect(302, "/")
		return
	}

	t, _ := time.Parse(time.RFC3339, topic.CreateTime)
	topic.CreateTime = t.Format("2006-01-02 15:04:05")

	log.Debugf("imgData:%+v", imgData)
	isTail := !(len(imgData) == num)
	g.HTML(http.StatusOK, "pic_detail.html", gin.H{
		"items":      topic,
		"imgs":       imgData,
		"page":       page,
		"next_page":  page + 1,
		"back_page":  page - 1,
		"imgs_count": len(imgData),
		"istail":     isTail,
	})
}

func TopicNice(g *gin.Context) {
	tid := g.Param("tid")
	state, _ := strconv.Atoi(g.Param("state"))
	page := g.Param("page")

	if len(tid) == 0 {
		g.Redirect(404, "/404")
		return
	}
	store.SetTopicNice(tid, state)
	g.Redirect(302, fmt.Sprintf("/d/%s/%s", tid, page))
}
