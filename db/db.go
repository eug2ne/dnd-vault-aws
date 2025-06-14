package db

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UserData struct {
	ID       string `dynamodbav:"id"`
	Username string `dynamodbav:"username"`
	Password string `dynamodbav:"password"`
	Usertype string `dynamodbav:"usertype"`
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

func (user UserData) GetKey() map[string]types.AttributeValue {
	// return primary keys in format that can be sent to db
	id, err := attributevalue.Marshal(user)
	if err != nil {
		panic(err)
	}
	username, err := attributevalue.Marshal(user.Username)
	if err != nil {
		panic(err)
	}
	usertype, err := attributevalue.Marshal(user.Usertype)
	if err != nil {
		panic(err)
	}

	return map[string]types.AttributeValue{"id": id, "username": username, "usertype": usertype}
}

func (user UserData) String() string {
	// return string format of user data
	return fmt.Sprintf("ID: %v\n\tUsername: %v\n\tPassword: %v\n\tUsertype:%v\n",
		user.ID, user.Username, user.Password, user.Usertype)
}
