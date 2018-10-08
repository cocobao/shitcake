package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/cocobao/log"
	"github.com/cocobao/shitcake/store"
	"github.com/gin-gonic/gin"
)

func Home(g *gin.Context) {
	var page int
	if v := g.Param("page"); len(v) > 0 {
		page, _ = strconv.Atoi(v)

	} else {
		page = 0
	}

	var num int
	if v := g.Param("num"); len(v) > 0 {
		num, _ = strconv.Atoi(v)
		if num > 10 {
			num = 10
		}
	} else {
		num = 10
	}

	topic, _ := store.GetTopics(num, page)

	for _, val := range topic {
		t, err := time.Parse(time.RFC3339, val.CreateTime)
		if err != nil {
			log.Warnf("parse time fail,err:%v", err)
		}
		st := t.Format("2006-01-02 15:04:05")
		val.CreateTime = st

		log.Debugf("st:%s, %#v", st, val)
	}

	g.HTML(http.StatusOK, "home.html", gin.H{
		"items": topic,
		"page":  page,
		"num":   num,
	})
}
