package controller

import (
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/Unknwon/com"
	"github.com/cocobao/shitcake/model"
	"github.com/gin-gonic/gin"

	log "github.com/cihub/seelog"
	"github.com/cocobao/shitcake/store"
	"github.com/cocobao/shitcake/utils"
)

type UploadController struct {
	BaseController
}

func NewUploadController(c *gin.Context) *UploadController {
	ctrl := &UploadController{}
	ctrl.ginCtx = c
	return ctrl
}

func (c *UploadController) Get() {
	c.TurnToPage("uploadImage.html")
}

func (c *UploadController) Post() {
	topicId := time.Now().Unix()

	var err error
	err = c.SaveIcon(topicId)
	if err != nil {
		log.Warn("save icon fail", err)
		c.ginCtx.Redirect(301, "/upload")
		return
	}
	err = c.staticImageList(topicId)
	if err != nil {
		log.Warn("save images fail", err)
		c.ginCtx.Redirect(301, "/upload")
		return
	}

	title := c.ginCtx.PostForm("title")
	isVip := c.ginCtx.PostForm("isVip")
	category, _ := strconv.Atoi(c.ginCtx.PostForm("category"))

	store.Db.SaveImageTopic(&model.ImageTopic{
		TopicID:    topicId,
		Title:      title,
		Category:   category,
		IsVip:      isVip,
		CreateTime: utils.TimeSecToString(time.Now().Unix()),
		SeeTime:    0,
	})
	c.TurnToPage("uploadImage.html")
}

func (c *UploadController) SaveIcon(topicId int64) error {
	fromFile, header, err := c.ginCtx.Request.FormFile("icon")
	if err != nil {
		log.Warn("get icon file fail", err)
	}
	defer fromFile.Close()

	pt := path.Join("static/icon", fmt.Sprintf("/%d/%s", topicId, header.Filename))
	log.Debug("img:", pt)
	if !com.IsExist(pt) {
		os.MkdirAll(path.Dir(pt), os.ModePerm)
	}

	//生成目标本地文件
	var dst *os.File
	dst, err = os.Create(pt)
	defer dst.Close()
	if err != nil {
		log.Warn("dst create fail", err)
		return err
	}

	//拷贝图片数据到本地文件
	if _, err = io.Copy(dst, fromFile); err != nil {
		log.Warn("io copy fail", err)
		return err
	}
	return nil
}

func (c *UploadController) staticImageList(topicId int64) error {
	from, err := c.ginCtx.MultipartForm()
	if err != nil {
		log.Error("get MultipartForm fail", err)
		return err
	}

	files := from.File["multiImages"]
	if files == nil || len(files) == 0 {
		fmt.Errorf("no multi file found,%v", files)
	}
	for _, oneFile := range files {
		pt := path.Join("static/images", fmt.Sprintf("%d/%s", topicId, oneFile.Filename))
		log.Debug("img:", pt)
		if !com.IsExist(pt) {
			os.MkdirAll(path.Dir(pt), os.ModePerm)
		}
		if err := c.ginCtx.SaveUploadedFile(oneFile, pt); err != nil {
			return fmt.Errorf("save file fail,%d, %s,%v", topicId, oneFile.Filename, err)
		}
		log.Debug("save img success,", oneFile.Filename)
	}

	return nil
}
