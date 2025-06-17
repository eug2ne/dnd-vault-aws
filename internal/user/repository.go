// / db operations related to user
package user

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gin-gonic/gin"
)

type Repository struct {
	DB        *dynamodb.Client
	TableName string
}

func NewRepository(db *dynamodb.Client) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) AddUser(ctx *gin.Context, user UserData) error {
	// add new user data to db
	item, err := attributevalue.MarshalMap(user)
	if err != nil {
		panic(err)
	}
	_, err = r.DB.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(r.TableName),
		Item:      item,
	})
	if err != nil {
		log.Printf("Couldn't add item to table. Here's why: %v\n", err)
	}
	return err
}

func (r *Repository) UpdateUser(ctx *gin.Context, user UserData) (map[string]map[string]interface{}, error) {
	// update user data in db
	var err error
	var response *dynamodb.UpdateItemOutput
	var attributeMap map[string]map[string]interface{}

	update := expression.Set(expression.Name("password"), expression.Value(user.Password))
	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		log.Printf("Couldn't build expression for update. Here's why: %v\n", err)
	} else {
		response, err = r.DB.UpdateItem(ctx, &dynamodb.UpdateItemInput{
			TableName:                 aws.String(r.TableName),
			Key:                       user.GetKey(),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			UpdateExpression:          expr.Update(),
			ReturnValues:              types.ReturnValueUpdatedNew,
		})
		if err != nil {
			log.Printf("Couldn't update user %v. Here's why: %v\n", user.UserName, err)
		} else {
			err = attributevalue.UnmarshalMap(response.Attributes, &attributeMap)
			if err != nil {
				log.Printf("Couldn't unmarshall update response. Here's why: %v\n", err)
			}
		}
	}
	return attributeMap, err
}

func (r *Repository) GetUserByID(ctx *gin.Context, id string) (*UserData, error) {
	user := UserData{UserID: id}
	response, err := r.DB.GetItem(ctx, &dynamodb.GetItemInput{
		Key:       user.GetKey(),
		TableName: aws.String(r.TableName),
	})
	if err != nil {
		log.Printf("Couldn't get info about %v. Here's why: %v\n", user.UserName, err)
	} else {
		err = attributevalue.UnmarshalMap(response.Item, &user)
		if err != nil {
			log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
		}
	}
	return &user, err
}

func (r *Repository) GetUserbyEmail(ctx *gin.Context, email string) (*UserData, error) {
	gsiPK := "EMAIL#" + email

	input := &dynamodb.QueryInput{
		TableName:              aws.String(r.TableName),
		IndexName:              aws.String("GSI1"),
		KeyConditionExpression: aws.String("GSIPK = :gsiPK"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":gsiPK": &types.AttributeValueMemberS{Value: gsiPK},
		},
		Limit: aws.Int32(1),
	}

	result, err := r.DB.Query(ctx, input)
	if err != nil {
		return nil, err
	}

	if len(result.Items) == 0 {
		return nil, fmt.Errorf("user not found")
	}

	var user UserData
	err = attributevalue.UnmarshalMap(result.Items[0], &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) DeleteUser(ctx context.Context, user UserData) error {
	_, err := r.DB.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(r.TableName),
		Key:       user.GetKey(),
	})
	if err != nil {
		log.Printf("Couldn't delete %v from the table. Here's why: %v\n", user.UserName, err)
	}
	return err
}
