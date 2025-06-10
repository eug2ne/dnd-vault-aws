package api

import (
	"net/http"
	"user/vault/auth"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	if err := auth.Authorize(c); err != nil {
		er := http.StatusUnauthorized
		http.Error(c.Writer, "Authorization failed: User has no authorization", er)
		return
	}

	// clear cookie
	c.SetCookie("session_token", "", -1, "/", "localhost:8080", false, true)
	c.SetCookie("csrf_token", "", -1, "/", "localhost:8080", false, false)

	// clear token from Users
	username := c.Request.FormValue("username")
	auth.Users[username] = auth.Login{
		HashedPassword: "",
		SessionToken:   "",
		CSRFToken:      "",
	}

	// return 200
	c.IndentedJSON(http.StatusOK, auth.Users[username])
}
