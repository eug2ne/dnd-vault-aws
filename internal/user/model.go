package user

import (
	"fmt"
	"user/vault/internal/group"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UserData struct {
	PK       string             `dynamodbav:"PK"` // USER#(UserID)
	SK       string             `dynamodbav:"SK"` // PROFILE
	UserID   string             `dynamodbav:"userid"`
	UserName string             `dynamodbav:"username"`
	Password string             `dynamodbav:"password"`
	UserType string             `dynamodbav:"usertype"`
	Groups   []group.MemberData `dynamodbav:"groups"`
	Email    string             `dynamodbav:"email"`
	Session  SessionData        `dynamodbav:"session"`
}

type SessionData struct {
	SessionToken string `dynamodbav:"sessiontoken"`
	CSRFToken    string `dynamodbav:"csrftoken"`
}

func (user UserData) GetKey() map[string]types.AttributeValue {
	// return PK, SK, UserID in format that can be sent to db
	id, err := attributevalue.Marshal(user.UserID)
	if err != nil {
		panic(err)
	}
	pk, err := attributevalue.Marshal(user.PK)
	if err != nil {
		panic(err)
	}
	sk, err := attributevalue.Marshal(user.SK)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"PK": pk, "SK": sk, "id": id}
}

func (user UserData) String() string {
	// return string format of user data
	return fmt.Sprintf("ID: %v\n\tUserName: %v\n\tPassword: %v\n\tUserType:%v\n, Email: %v\n",
		user.UserID, user.UserName, user.Password, user.UserType, user.Email)
}

func (user UserData) GetGroups() string {
	// TODO: return group data from user data
	return fmt.Sprintf("User %v is currently joining these groups:\n",
		user.UserName)
}
