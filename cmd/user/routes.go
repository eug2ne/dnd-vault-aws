package user

import (
	"user/vault/internal/user"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, db *dynamodb.Client) {
	// initialize user repo + handler
	userRepo := user.NewRepository(db)
	userHandler := user.NewHandler(userRepo)
	// add api paths
	router.POST("/signup", userHandler.SignUp)
	router.GET("/login", userHandler.Login)
	router.GET("/:id/profile", userHandler.GetProfile)
	router.POST("/:id/profile", userHandler.UpdateProfile)
	router.GET("/:id/groups", userHandler.GetGroups)
}
