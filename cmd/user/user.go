package user

import (
	"net/http"
	"user/vault/auth"

	"github.com/gin-gonic/gin"
)

func FetchUserData(c *gin.Context) {
	if auth.SearchAuthError(c) {
		// user auth failed
		// TODO: return auth fail/404 error html page
		c.IndentedJSON(http.StatusUnauthorized, "Can't access User page. You need to login to access this page.")
		return
	}

	if t, ok := c.Params.Get("usertype"); ok {
		if t == "dm" {
			// fetch dm data
			// return all player+dm datas

		} else if t == "player" {
			// fetch player data
			// return only user data
		} else {
			// invaild user type
			er := http.StatusBadRequest
			http.Error(c.Writer, "Invalid user type", er)
			return
		}
	}
	c.IndentedJSON(http.StatusOK, auth.SearchAuthError(c))
}
