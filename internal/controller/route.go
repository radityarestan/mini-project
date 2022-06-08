package controller

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(ctr *Controller) *gin.Engine {
	router := gin.Default()

	router.POST("/api/users/register", ctr.register)
	router.GET("/api/users/login", ctr.login)

	return router
}
