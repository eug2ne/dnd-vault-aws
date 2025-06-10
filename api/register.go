package api

import (
	"net/http"
	"user/vault/auth"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	// register individual password for each user
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	usertype := c.Request.FormValue("usertype")
	if len(password) < 8 {
		er := http.StatusNotAcceptable
		http.Error(c.Writer, "Invalid password", er)
		return
	}

	// check if user data exist in db
	userdata := auth.FindUserData(username)
	if userdata.Usertype == "nonexistant" {
		// user data do not exist in db
		if usertype == "dm" {
			// user trying to add another dm
			// send error
			er := http.StatusConflict
			http.Error(c.Writer, "Duplicate DM: There can be only 1 DM!", er)
			return
		}
		// send error
		er := http.StatusNotAcceptable
		http.Error(c.Writer, "Player does not exist in db: Who are you?!", er)
		return
	} else {
		// user already existing in db
		// create hashed password
		hashedPassword, _ := auth.HashPassword(password)
		// change password of user data in db
		userdata.Password = hashedPassword

		// return 200
		c.IndentedJSON(http.StatusOK, userdata)
	}
}
