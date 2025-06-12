package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup) {
	// add api paths
	router.POST("/", FetchUserData)
}
