package routers

import (
	"github.com/cocobao/shitcake/controller"
	"github.com/gin-gonic/gin"
)

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

func TopicDetail(c *gin.Context) {
	controller.NewTopicController(c).Get()
}

func DeleteTopic(c *gin.Context) {
	controller.NewTopicController(c).Del()
}

func InsertTopic(c *gin.Context) {
	controller.NewUploadController(c).SimpleInsertTopic()
}
