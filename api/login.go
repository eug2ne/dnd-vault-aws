package api

import (
	"fmt"
	"net/http"
	"user/vault/auth"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login api")

	if r.Method != http.MethodPost {
		er := http.StatusMethodNotAllowed
		http.Error(w, "Invalid method", er)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	if len(password) < 8 {
		er := http.StatusNotAcceptable
		http.Error(w, "Invalid password", er)
		return
	}

	if username == "dungeon_master" {
		// dm auth
		// check hashedPassword
		hashedPassword, _ := auth.HashPassword(password)
		if hashedPassword != auth.Users["dungeon_master"].HashedPassword {
			er := http.StatusConflict
			http.Error(w, "Wrong password", er)
			return
		}
	} else {
		// player auth
		// check player exist in users
		if _, ok := auth.Users[username]; !ok {
			er := http.StatusConflict
			http.Error(w, "Player not exist", er)
			return
		}
	}

}
