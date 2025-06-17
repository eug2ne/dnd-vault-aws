package group

import "github.com/gin-gonic/gin"

type Handler struct {
	Repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{Repo: repo}
}

func (handler Handler) CreateNewGroup(c *gin.Context) {
	// TODO: create new group
}

func (handler Handler) GetGroupData(c *gin.Context) {
	// TODO: retrieve group data from db
}

func (handler Handler) GetMembers(c *gin.Context) {
	// TODO: retrieve members data from db
}
