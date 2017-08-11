package routers

import (
	"github.com/cocobao/shitcake/controller"
	"github.com/gin-gonic/gin"
)

func HomeGet(c *gin.Context) {
	controller.NewHomeController(c).Get()
}

func LoginGet(c *gin.Context) {
	controller.NewLogController(c).Get()
}

func LoginPost(c *gin.Context) {
	controller.NewLogController(c).Post()
}

func UploadGet(c *gin.Context) {
	controller.NewUploadController(c).Get()
}

func UploadPost(c *gin.Context) {
	controller.NewUploadController(c).Post()
}
