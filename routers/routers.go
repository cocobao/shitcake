package routers

import (
	"github.com/cocobao/shitcake/conf"
	"github.com/cocobao/shitcake/controller"
	"github.com/gin-gonic/gin"
)

func LoadRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(controller.Auth())
	router.LoadHTMLGlob(conf.GCfg.ViewPath)
	router.Static("/static", conf.GCfg.StaticPath)
	router.GET("/", HomeGet)
	router.GET("/login", LoginGet)
	router.POST("/login", LoginPost)
	router.GET("/upload", UploadGet)
	router.POST("/upload", UploadPost)
	router.GET("/topicdetail", TopicDetail)
	router.GET("/deltopic", DeleteTopic)
	return router
}

func Run() {
	engin := LoadRouter()
	engin.Run()
}
