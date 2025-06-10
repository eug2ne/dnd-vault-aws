package auth

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

var AuthError error = errors.New("Authorization failed")

func Authorize(c *gin.Context) error {
	username := c.Request.FormValue("username")
	userdata := FindUserData(username)
	if userdata.Usertype == "nonexistant" {
		return AuthError
	}

	// TODO: (BUG) authorize producing AuthError even when including csrf token in header
	// get session token from cookie
	st, err := c.Request.Cookie("session_token")
	if err != nil || st.Value == "" || st.Value != Users[username].SessionToken {
		return AuthError
	}

	// get csrf token from header
	csrf := c.Request.Header.Get("X-CSRF-TOKEN")
	fmt.Println(csrf)
	if csrf != Users[username].CSRFToken || csrf == "" {
		return AuthError
	}

	return nil
}
