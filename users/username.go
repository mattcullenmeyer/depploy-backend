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
	UserId   string
	Username string
}

func Profile(c *gin.Context) {
	username := c.Params.ByName("username")

	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)

	result, err := svc.Query(&dynamodb.QueryInput{
		TableName: aws.String(tableName),
		Limit:     aws.Int64(1),
		KeyConditions: map[string]*dynamodb.Condition{
			"Username": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(username),
					},
				},
			},
		},
		IndexName: aws.String("Username-index"),
	})
	if err != nil {
		log.Fatalf("Got error calling query: %s", err)
	}

	users := []User{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &users)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	if len(users) == 0 {
		c.AbortWithStatus((http.StatusNotFound))
	}

	user := users[0]

	c.JSON(http.StatusOK, user.UserId)
}
