package routers

import (
	"github.com/cocobao/shitcake/controller"
	"github.com/gin-gonic/gin"
)

func LoadRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(controller.Auth())
	// router.LoadHTMLGlob("views/learn/*")
	router.LoadHTMLGlob("views/*.html")
	router.Static("/static", "./static")
	router.Static("/static2", "/Users/ybz/Documents/static")
	router.GET("/", HomeGet)
	router.GET("/login", LoginGet)
	router.POST("/login", LoginPost)
	router.GET("/upload", UploadGet)
	router.POST("/upload", UploadPost)
	router.GET("/topicdetail", TopicDetail)
	router.GET("/deltopic", DeleteTopic)
	router.POST("/insert-topic", InsertTopic)
	return router
}

func Run() {
	engin := LoadRouter()
	engin.Run()
}
