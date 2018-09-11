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

	"github.com/cocobao/log"
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
	title := c.ginCtx.PostForm("title")
	isVip := c.ginCtx.PostForm("isVip")
	category, _ := strconv.Atoi(c.ginCtx.PostForm("category"))
	topicID := strconv.Itoa(int(time.Now().Unix()))

	topic := &model.ImageTopic{
		TopicID:    topicID,
		Title:      title,
		Category:   category,
		IsVip:      isVip,
		CreateTime: utils.TimeSecToString(time.Now().Unix()),
		SeeTime:    0,
	}

	var err error
	err = c.SaveIcon(topic)
	if err != nil {
		log.Warn("save icon fail", err)
		c.ginCtx.Redirect(301, "/upload")
		return
	}
	err = c.staticImageList(topic)
	if err != nil {
		log.Warn("save images fail", err)
		c.ginCtx.Redirect(301, "/upload")
		return
	}

	store.Db.SaveImageTopic(topic)
	c.TurnToPage("uploadImage.html")
}

func (c *UploadController) SaveIcon(model *model.ImageTopic) error {
	fromFile, _, err := c.ginCtx.Request.FormFile("icon")
	if err != nil {
		log.Warn("get icon file fail", err)
	}
	defer fromFile.Close()

	ranName := utils.Md5StringByNowTime()
	// u, _ := url.Parse(header.Filename)
	// model.Icon = u.EscapedPath()
	model.Icon = ranName

	pt := path.Join("static/icon", fmt.Sprintf("/%s/%s", model.TopicID, ranName))
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

func (c *UploadController) staticImageList(model *model.ImageTopic) error {
	from, err := c.ginCtx.MultipartForm()
	if err != nil {
		log.Error("get MultipartForm fail", err)
		return err
	}

	files := from.File["multiImages"]
	if files == nil || len(files) == 0 {
		return fmt.Errorf("no multi file found,%v", files)
	}
	var fileNames []string
	for _, oneFile := range files {
		// u, _ := url.Parse(oneFile.Filename)
		// name := u.EscapedPath()
		ranName := utils.Md5StringByNowTime()

		pt := path.Join("static/images", fmt.Sprintf("%s/%s", model.TopicID, ranName))
		if !com.IsExist(pt) {
			os.MkdirAll(path.Dir(pt), os.ModePerm)
		}
		fileNames = append(fileNames, ranName)
		if err := c.ginCtx.SaveUploadedFile(oneFile, pt); err != nil {
			return fmt.Errorf("save file fail,%s, %s,%v", model.TopicID, ranName, err)
		}
		log.Debug("save img success,", oneFile.Filename)
	}
	model.Images = fileNames

	return nil
}
