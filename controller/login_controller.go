package controller

import (
	"github.com/gin-gonic/gin"
)

var IsLogin bool

type LogController struct {
	BaseController
}

func NewLogController(c *gin.Context) *LogController {
	ctrl := &LogController{}
	ctrl.ginCtx = c
	return ctrl
}

func (c *LogController) Get() {
	c.TurnToPage("login.html")
}

func (c *LogController) Post() {
	uname := c.ginCtx.Request.PostFormValue("uname")
	pwd := c.ginCtx.Request.PostFormValue("pwd")

	if uname != "admin" || pwd != "admin" {
		c.ginCtx.Redirect(301, "/login")
	}
	IsLogin = true
	c.TurnToPage("uploadImage.html")
}
