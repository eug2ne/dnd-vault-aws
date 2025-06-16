package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup) {
	// TODO: initialize user repo + handler
	// add api paths
	router.POST("/", FetchUserData)
}
