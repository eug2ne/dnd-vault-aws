package api

import (
	"fmt"
	"net/http"
	"user/vault/auth"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	if len(password) < 8 {
		er := http.StatusNotAcceptable
		http.Error(c.Writer, "Invalid password", er)
		return
	}

	if username == "dungeon_master" {
		// dm signup
		// add new user
		hashedPassword, _ := auth.HashPassword(password)
		if hashedPassword != auth.Users["dungeon_master"].HashedPassword {
			er := http.StatusConflict
			http.Error(c.Writer, "Wrong password", er)
			return
		}
	} else {
		// player auth
		// check player exist in users
		if _, ok := auth.Users[username]; !ok {
			er := http.StatusConflict
			http.Error(c.Writer, "Player not exist", er)
			return
		}
	}
}

func Login(c *gin.Context) {
	fmt.Println("login api")

	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	if len(password) < 8 {
		er := http.StatusNotAcceptable
		http.Error(c.Writer, "Invalid password", er)
		return
	}

	if username == "dungeon_master" {
		// dm auth
		// check hashedPassword
		hashedPassword, _ := auth.HashPassword(password)
		if hashedPassword != auth.Users["dungeon_master"].HashedPassword {
			er := http.StatusConflict
			http.Error(c.Writer, "Wrong password", er)
			return
		}
	} else {
		// player auth
		// check player exist in users
		if _, ok := auth.Users[username]; !ok {
			er := http.StatusConflict
			http.Error(c.Writer, "Player not exist", er)
			return
		}
	}
}
