package auth

type Login struct {
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}

// TODO: create Users into JSON
var Users = map[string]Login{}
