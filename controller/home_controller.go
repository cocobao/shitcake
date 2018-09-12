package controller

import (
	"net/http"

	"github.com/cocobao/shitcake/store"
	"github.com/cocobao/shitcake/utils"
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
	topic := store.Db.GetImageTopic(10)

	var i []interface{}
	for _, v := range topic {
		m := utils.StructToMapJson(v)
		delete(m, "images")
		i = append(i, m)
	}

	c.ginCtx.HTML(http.StatusOK, "home.html", gin.H{
		"files": i,
	})
}
