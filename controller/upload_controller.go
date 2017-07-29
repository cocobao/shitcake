package controller

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
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
	if !IsLogin {
		c.ginCtx.Redirect(301, "/login")
		return
	}
	c.TurnToPage("uploadImage.html")
}

func (c *UploadController) Post() {
	// multiImage, err := c.GetFiles("multiImages")
	// if err != nil || multiImage == nil {
	// 	c.gotError("get multiImage err", err)
	// 	return
	// }
	// topicId := time.Now().Unix()
	// c.staticImageList(multiImage, topicId)
}

func (c *UploadController) staticImageList(multiImage []*multipart.FileHeader, topicId int64) (imgList []string, err error) {
	// finish := true
	// for i, fh := range multiImage {
	// 	randName := utils.Md5StringByNowTime()

	// 	func() {
	// 		file, err := fh.Open()
	// 		if err != nil {
	// 			beego.Error("open multiImage fail", i, err)
	// 			finish = false
	// 			return
	// 		}
	// 		defer file.Close()

	// 		pt := path.Join("static", fmt.Sprintf("images/%d/%s", topicId, randName))
	// 		if !com.IsExist(pt) {
	// 			os.MkdirAll(path.Dir(pt), os.ModePerm)
	// 		}

	// 		//生成目标本地文件
	// 		var dst *os.File
	// 		dst, err = os.Create(pt)
	// 		defer dst.Close()
	// 		if err != nil {
	// 			beego.Error("dst create fail", i, err)
	// 			return
	// 		}

	// 		//拷贝图片数据到本地文件
	// 		if _, err = io.Copy(dst, file); err != nil {
	// 			beego.Error("io copy fail", i, err)
	// 			return
	// 		}
	// 	}()
	// }

	// if !finish {

	// }
	return
}
