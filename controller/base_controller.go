package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
	ginCtx *gin.Context
}

func (c *BaseController) TurnToPage(page string) {
	c.ginCtx.HTML(http.StatusOK, page, gin.H{})
}
