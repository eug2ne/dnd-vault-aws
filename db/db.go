package db

type UserData struct {
	Username string
	Password string
	Usertype string
}

// temp db (default value)
var DB = []UserData{
	{Username: "player1",
		Password: "letsplaydnd",
		Usertype: "player"},
	{Username: "player2",
		Password: "letsplaydnd2",
		Usertype: "player"},
	{Username: "player3",
		Password: "letsplaydnd",
		Usertype: "player"},
	{Username: "player4",
		Password: "letsplaydnd",
		Usertype: "player"},
	{Username: "player5",
		Password: "letsplaydnd",
		Usertype: "player"},
	{Username: "player6",
		Password: "letsplaydnd",
		Usertype: "player"},
	{Username: "dungeon_master1",
		Password: "iamdungeonmaster",
		Usertype: "dm"},
}
