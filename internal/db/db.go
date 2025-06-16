package db

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewDynamoDBClient(config aws.Config) *dynamodb.Client {
	// create + return new db client
	return dynamodb.NewFromConfig(config)
}
