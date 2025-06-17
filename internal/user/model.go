package user

import (
	"fmt"
	"user/vault/internal/group"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UserData struct {
	PK       string             `dynamodbav:"PK"`     // USER#(UserID)
	SK       string             `dynamodbav:"SK"`     // METADATA
	GSI1PK   string             `dynamodbav:"GSI1PK"` // EMAIL#(Email)
	UserID   string             `dynamodbav:"userid"`
	UserName string             `dynamodbav:"username"`
	Password string             `dynamodbav:"password"`
	UserType string             `dynamodbav:"usertype"`
	Groups   []group.MemberData `dynamodbav:"groups"`
	Email    string             `dynamodbav:"email"`
	Session  SessionData        `dynamodbav:"session"`
}

type SessionData struct {
	SessionToken string `dynamodbav:"sessiontoken" json:"sessiontoken"`
	CSRFToken    string `dynamodbav:"csrftoken" json:"csrftoken"`
}

type SignupRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Type     string `json:"type" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserDataRequest struct {
	Name     string      `json:"name" binding:"required"`
	Email    string      `json:"email" binding:"required,email"`
	Type     string      `json:"type" binding:"required"`
	Password string      `json:"password" binding:"required"`
	Session  SessionData `json:"session" binding:"required"`
}

func NewUserData(userID string, userName string, password string, userType string, email string) *UserData {
	return &UserData{
		PK:       "USER#" + userID,
		SK:       "METADATA",
		GSI1PK:   "EMAIL#" + email,
		UserID:   userID,
		UserName: userName,
		Password: password,
		UserType: userType,
		Groups:   []group.MemberData{},
		Email:    email,
		Session:  SessionData{},
	}
}

func (user UserData) GetKey() map[string]types.AttributeValue {
	// return UserID in format that can be sent to db
	id, err := attributevalue.Marshal(user.UserID)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"id": id}
}

func (user UserData) String() string {
	// return string format of user data
	return fmt.Sprintf("ID: %v\n\tUserName: %v\n\tPassword: %v\n\tUserType:%v\n, Email: %v\n",
		user.UserID, user.UserName, user.Password, user.UserType, user.Email)
}
