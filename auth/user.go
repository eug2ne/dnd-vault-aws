package auth

type Login struct {
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}

var Users = map[string]Login{}
