package group

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type GroupData struct {
	PK            string           `dynamodbav:"PK"` // GROUP#(GroupID)
	SK            string           `dynamodbav:"SK"` // METADATA
	GroupID       string           `dynamodbav:"groupid"`
	GroupName     string           `dynamodbav:"groupname"`
	DM            MemberData       `dynamodbav:"dm"`
	Players       []MemberData     `dynamodbav:"players"`
	Private       bool             `dynamodbav:"private"`
	GroupPassword string           `dynamodbav:"grouppassword"` // for private groups
	CreatedAt     string           `dynamodbav:"createdat"`
	Invitations   []InvitationData `dynamodbav:"invitations"`
}

type MemberData struct {
	PK       string `dynamodbav:"PK"` // GROUP#(GroupID)
	SK       string `dynamodbav:"SK"` // MEMBER#(UserID)
	UserID   string `dynamodbav:"userid"`
	GroupID  string `dynamodbav:"groupid"`
	DMing    bool   `dynamodbav:"dming"`
	JoinedAt string `dynamodbav:"joinedat"`
}

type InvitationData struct {
	PK      string `dynamodbav:"PK"` // GROUP#(GroupID)
	SK      string `dynamodbav:"SK"` // INVITE#(UserID)
	UserID  string `dynamodbav:"userid"`
	GroupID string `dynamodbav:"groupid"`
	Status  string `dynamodbav:"status"`
}

func (group GroupData) GetKey() types.AttributeValue {
	// return GroupID in format that can be sent to db
	id, err := attributevalue.Marshal(group.GroupID)
	if err != nil {
		panic(err)
	}
	return id
}

func (group GroupData) String() string {
	// return string format of group data
	return fmt.Sprintf("ID: %v\n\tGroupName: %v\n\tDM: %v\n\tPrivate: %v\n\tCreated: %v\n",
		group.GroupID, group.GroupName, group.DM.UserID, group.Private, group.CreatedAt)
}

func (group GroupData) GetMembers() string {
	// TODO: return member data from group data
	return fmt.Sprintf("Group %v has %v players and is dmed by %v",
		group.GroupName, len(group.Players), group.DM.UserID)
}
