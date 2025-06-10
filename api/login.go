package api

import (
	"fmt"
	"net/http"
	"time"
	"user/vault/auth"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	fmt.Println("login api")

	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	userdata := auth.FindUserData(username)
	// check uesr exist in db
	if userdata.Usertype == "nonexistant" {
		er := http.StatusUnauthorized
		http.Error(c.Writer, "Invalid player/DM", er)
		return
	}

	// check password
	if !auth.CheckPasswordHash(password, userdata.Password) {
		er := http.StatusUnauthorized
		http.Error(c.Writer, "Invalid password", er)
		return
	}

	// create session token + csrf token for user
	sessionToken := auth.CreateToken(32)
	csrfToken := auth.CreateToken(32)
	// set session + csrf cookie
	c.SetCookie("session_token", sessionToken, int(time.Hour), "/", "localhost:8080", false, true)
	c.SetCookie("csrf_token", csrfToken, int(time.Hour), "/", "localhost:8080", false, false)
	// store session + csrf token in Users
	auth.Users[username] = auth.Login{
		SessionToken: sessionToken,
		CSRFToken:    csrfToken,
	}

	// return 200
	c.IndentedJSON(http.StatusOK, auth.Users[username])
}
