package db

type UserData struct {
	Username string
	Password string
	Usertype string
}

// temp db (default value)
var DB = []UserData{
	{Username: "player1",
		Password: "letsplaydnd1",
		Usertype: "player"},
	{Username: "player2",
		Password: "letsplaydnd2",
		Usertype: "player"},
	{Username: "player3",
		Password: "letsplaydnd1",
		Usertype: "player"},
	{Username: "player4",
		Password: "letsplaydnd1",
		Usertype: "player"},
	{Username: "dungeon_master1",
		Password: "dmingthegame1",
		Usertype: "dm"},
}
