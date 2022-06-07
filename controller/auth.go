package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctr *Controller) register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func (ctr *Controller) login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}
