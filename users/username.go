package users

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string
}

func Profile(c *gin.Context) {
	username := c.Params.ByName("username")

	tableName := os.Getenv("DYNAMODB_TABLE_NAME")
	endpoint := os.Getenv("DYNAMODB_ENDPOINT")

	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String(endpoint), // uses default generated endpoint if an empty string
	}))
	svc := dynamodb.New(sess)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Username": {
				S: aws.String(username),
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if result.Item == nil {
		c.JSON(http.StatusNotFound, "user not found")
		return
	}

	user := User{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	c.JSON(http.StatusOK, user.Username)
}
