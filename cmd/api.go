package cmd

import (
	"database/sql"
	"log"
	"net/http"
	"runtime/debug"
	"user/vault/api"

	"github.com/gin-gonic/gin"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func recoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// defer func to catch any panics
		defer func() {
			if err := recover(); err != nil {
				// log panic + stack trace
				msg := "Caught panic: %v, Stack trace: %s"
				log.Print(msg, err, string(debug.Stack()))

				// return generic error message to client
				er := http.StatusInternalServerError
				http.Error(c.Writer, "Internal Server Error", er)
			}
		}()
	}
}

func NewAPIServer(new_addr string, new_db *sql.DB) *APIServer {
	return &APIServer{addr: new_addr, db: new_db}
}

func (s *APIServer) Run() error {
	// create router
	router := gin.Default()
	router.Use(recoveryMiddleware())
	// create subrouter
	api_router := router.Group("/api")

	// add handlers tp subrouter
	api.RegisterRoutes(api_router)

	return router.Run(":8080")
}
