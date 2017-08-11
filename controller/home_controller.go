package controller

import (
	"net/http"

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
	topic := store.Db.GetImageTopic(10)

	var images []interface{}
	images = append(images, topic)
	c.ginCtx.HTML(http.StatusOK, "home.html", gin.H{
		"images": images,
	})
}
