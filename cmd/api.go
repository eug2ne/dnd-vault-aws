package cmd

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"runtime/debug"
	"user/vault/api"
	"user/vault/auth"
	"user/vault/cmd/user"

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

func errorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		SimpleError := errors.New("Very Simple Error")
		c.Error(SimpleError)
	}
}

func authorizeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// run authorize func
		var err error
		defer func() {
			err = auth.Authorize(c)

			if err != nil {
				// authorization failure
				er := http.StatusUnauthorized
				http.Error(c.Writer, "Authorization failed: User has no authorization to access vault data", er)
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
	user_router := router.Group("/user")
	bot_router := router.Group("/bots")

	// add middleware to subrouter
	user_router.Use(authorizeMiddleware(), errorMiddleware())
	bot_router.Use(authorizeMiddleware())

	// add handlers tp subrouter
	api.RegisterRoutes(api_router)
	user.RegisterRoutes(user_router)

	// TODO: need to set path + add handler to dm_router, player_router
	// TODO: or just set dm/player path to router?

	return router.Run(":8080")
}
