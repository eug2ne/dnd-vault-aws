package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	// add api paths
	router.POST("/login", Login)
	router.POST("/register", Register)
	router.POST("/logout", Logout)
}
