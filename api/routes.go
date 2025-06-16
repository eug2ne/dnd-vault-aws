package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	// add api paths
	router.POST("/login", Login)
	router.POST("/signup", SignUp)
	router.POST("/logout", Logout)
}
