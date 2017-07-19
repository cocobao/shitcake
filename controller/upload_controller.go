package controller

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"sexWeb/utility"

	"time"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego"
)

type Upload struct {
	BaseController
}

func (c *Upload) Get() {
	// if !IsLogin {
	// 	beego.Warn("not login")
	// 	c.Redirect("/login", 301)
	// 	return
	// }
	c.TplName = "uploadImage.html"
}

func (c *Upload) Upload() {
	// title := this.Input().Get("title")
	// isVip, _ := strconv.Atoi(this.Input().Get("isVip"))
	// topicType, _ := strconv.Atoi(this.Input().Get("topicType"))

	multiImage, err := c.GetFiles("multiImages")
	if err != nil || multiImage == nil {
		c.gotError("get multiImage err", err)
		return
	}
	topicId := time.Now().Unix()
	c.staticImageList(multiImage, topicId)
}

func (c *Upload) staticImageList(multiImage []*multipart.FileHeader, topicId int64) (imgList []string, err error) {
	finish := true
	for i, fh := range multiImage {
		randName := utility.Md5StringByNowTime()

		func() {
			file, err := fh.Open()
			if err != nil {
				beego.Error("open multiImage fail", i, err)
				finish = false
				return
			}
			defer file.Close()

			pt := path.Join("static", fmt.Sprintf("images/%d/%s", topicId, randName))
			if !com.IsExist(pt) {
				os.MkdirAll(path.Dir(pt), os.ModePerm)
			}

			//生成目标本地文件
			var dst *os.File
			dst, err = os.Create(pt)
			defer dst.Close()
			if err != nil {
				beego.Error("dst create fail", i, err)
				return
			}

			//拷贝图片数据到本地文件
			if _, err = io.Copy(dst, file); err != nil {
				beego.Error("io copy fail", i, err)
				return
			}
		}()
	}

	if !finish {

	}
	return
}

func (c *Upload) gotError(msg string, err error) {
	beego.Error(msg, err)
	c.Redirect("/uploadimage", 301)
}
