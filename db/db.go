package db

type User struct {
	username string
	password string
	usertype string
}

var DB = []User{
	{username: "player1",
		password: "letsplaydnd1",
		usertype: "player"},
	{username: "player2",
		password: "letsplaydnd2",
		usertype: "player"},
	{username: "player3",
		password: "letsplaydnd1",
		usertype: "player"},
	{username: "player4",
		password: "letsplaydnd1",
		usertype: "player"},
	{username: "dungeon_master1",
		password: "dmingthegame1",
		usertype: "player"}}
