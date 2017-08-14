package routers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cocobao/shitcake/conf"
	"github.com/cocobao/shitcake/controller"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
)

func LoadRouter() http.Handler {
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(controller.Auth())
	router.LoadHTMLGlob("views/*")
	router.Static("/static", "./static")
	router.GET("/", HomeGet)
	router.GET("/login", LoginGet)
	router.POST("/login", LoginPost)
	router.GET("/upload", UploadGet)
	router.POST("/upload", UploadPost)
	router.GET("topicdetail", TopicDetail)
	return router
}

func Run() {
	err := gracehttp.Serve(
		&http.Server{
			Addr:    conf.GCfg.Port,
			Handler: LoadRouter(),
		},
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
