package controller

import (
	"net/http"

	"github.com/cocobao/log"
	"github.com/gin-gonic/gin"
)

type BaseApiResponse struct {
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

type BaseController struct {
	ginCtx *gin.Context
}

func (c *BaseController) ResponseSuccess(rsp interface{}) {
	c.ginCtx.JSON(http.StatusOK, rsp)
}

func (c *BaseController) TurnToPage(page string) {
	c.ginCtx.HTML(http.StatusOK, page, gin.H{})
}

func (c *BaseController) ParsePayload(payload interface{}) error {
	if err := c.ginCtx.BindJSON(payload); err != nil {
		c.Response(http.StatusForbidden, "JSON参数错误", nil)
		return err
	}
	log.Debug("payload:", payload)

	return nil
}

func (c *BaseController) Response(statusCode int, msg string, result interface{}) {
	response := &BaseApiResponse{}
	response.Msg = msg
	response.Result = result
	log.Debugf("http response:%v", response)
	c.ginCtx.JSON(statusCode, response)
}
