package main

import (
	"errors"
	"log"
	"net/http"
	"runtime/debug"
	"user/vault/internal/auth"

	"github.com/gin-gonic/gin"
)

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
