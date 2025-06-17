package group

import (
	"user/vault/internal/group"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, db *dynamodb.Client) {
	// initialize group repo + handler
	groupRepo := group.NewRepository(db)
	groupHandler := group.NewHandler(groupRepo)
	// add api paths
	router.GET("/:id", groupHandler.GetGroupData)
	router.POST("/create", groupHandler.CreateNewGroup)
	router.GET("/:id/members", groupHandler.GetMembers)
}
