package auth

import (
	"user/vault/db"
)

type Login struct {
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}

// stores login user data
var Users = map[string]Login{}

func FindUserData(username string) *db.UserData {
	// check if user data exist in db
	var u_index int = -1
	for i := 0; i < len(db.DB); i++ {
		if db.DB[i].Username == username {
			u_index = i
		}
	}

	if u_index == -1 {
		// user data do not exist in db
		return &db.UserData{
			Username: "not existing",
			Password: "0000",
			Usertype: "nonexistant",
		}
	} else {
		// return user data
		return &db.DB[u_index]
	}
}
