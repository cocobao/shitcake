package controller

import (
	"net/http"
	"time"

	"github.com/cocobao/log"
	"github.com/cocobao/shitcake/store"
	"github.com/gin-gonic/gin"
)

type HomeController struct {
	BaseController
}

func NewHomeController(c *gin.Context) *HomeController {
	ctrl := &HomeController{}
	ctrl.ginCtx = c
	return ctrl
}

func (c *HomeController) Get() {
	topic := store.GetTopics(10, 0)

	for _, v := range topic {
		v.Images = []string{}

		t, err := time.Parse(time.RFC3339, v.CreateTime)
		if err != nil {
			log.Warnf("parse time fail,err:%v", err)
		}
		st := t.Format("2006-01-02 15:04:05")
		v.CreateTime = st

		log.Debugf("st:%s, %#v", st, v)
	}

	c.ginCtx.HTML(http.StatusOK, "home.html", gin.H{
		"items": topic,
	})
}
