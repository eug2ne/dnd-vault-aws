package auth

import (
	"errors"
	"net/url"

	"github.com/gin-gonic/gin"
)

var AuthError error = errors.New("Authorization failed")

func Authorize(c *gin.Context) error {
	username := c.Request.FormValue("username")
	userdata := FindUserData(username)
	if userdata.Usertype == "nonexistant" {
		return c.Error(AuthError)
	}

	// get session token from cookie
	st, err := c.Request.Cookie("session_token")
	session, _ := url.QueryUnescape(st.Value)
	if err != nil || session == "" || session != Users[username].SessionToken {
		return c.Error(AuthError)
	}

	// get csrf token from header
	ct := c.Request.Header.Get("X-CSRF-TOKEN")
	csrf, _ := url.QueryUnescape(ct)
	if csrf != Users[username].CSRFToken || csrf == "" {
		return c.Error(AuthError)
	}

	return nil
}

func SearchAuthError(c *gin.Context) bool {
	for i := 0; i < len(c.Errors); i++ {
		if c.Errors[i].Error() == "Authorization failed" {
			return true
		}
	}

	return false
}
