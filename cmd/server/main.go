package main

import (
	"context"
	"log"
	"user/vault/api"
	"user/vault/cmd/user"
	"user/vault/internal/db"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	addr      string
	db_client *dynamodb.Client
}

func NewAPIServer(new_addr string, new_db *dynamodb.Client) *APIServer {
	return &APIServer{addr: new_addr, db_client: new_db}
}

func (s *APIServer) Run() error {
	// create router
	router := gin.Default()
	router.Use(recoveryMiddleware())
	// create subrouter
	api_router := router.Group("/api")
	user_router := router.Group("/user")
	bot_router := router.Group("/bots")

	// add middleware to subrouter
	user_router.Use(authorizeMiddleware(), errorMiddleware())
	bot_router.Use(authorizeMiddleware())

	// add handlers tp subrouter
	api.RegisterRoutes(api_router)
	user.RegisterRoutes(user_router, s.db_client)

	// TODO: need to set path + add handler to dm_router, player_router
	// TODO: or just set dm/player path to router?

	return router.Run(":8080")
}

func main() {
	ctx := context.TODO()

	// import AWS SDK
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// create Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	// get the first page of results for ListObjectsV2 for bucket
	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("dnd-vault-bucket"),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("first page results")
	for _, object := range output.Contents {
		log.Printf("key=%s size=%d", aws.ToString(object.Key), *object.Size)
	}

	// create db client
	dbClient := db.NewDynamoDBClient(cfg)

	// create + run new APIServer
	server := NewAPIServer(":8080", dbClient)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
