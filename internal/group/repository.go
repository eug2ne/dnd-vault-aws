// / db operations related to group
package group

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type Repository struct {
	DB        *dynamodb.Client
	TableName string
}

func NewRepository(db *dynamodb.Client) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) AddGroup(new_group GroupData) {
	// TODO: add new group data to db
}

func (r *Repository) UpdateGroup(update_group GroupData) {
	// TODO: update group data in db
}

func (r *Repository) DeleteGroup(delete_group GroupData) {
	// TODO: delete group data from db
}

func (r *Repository) AddMember(group GroupData, user_id string) {
	// TODO: create new member instance
	// TODO: add member instance to group data
	// TODO: get + update user data
}

func (r *Repository) DeleteMember(group GroupData, user_id string) {
	// TODO: get + delete user data
}

func (r *Repository) InviteMember(group GroupData, user_id string) {
	// TODO: create new invitation instance
	// TODO: add invitation instance to group data
}

func (r *Repository) UpdateInvitation(group GroupData, user_id string) {
	// TODO: get + update invitation data
	// if invitation.status == accepted
	r.AddMember(group, user_id)
	// if invitation.status == declined/expired
	r.DeleteInvitation(group, user_id)
}

func (r *Repository) DeleteInvitation(group GroupData, user_id string) {
	// TODO: delete invitation data
}
