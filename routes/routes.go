package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
)

type User struct {
	UserId string
	Username string
}

func username(c *gin.Context) {
	username := c.Params.ByName("username")
	fmt.Println(username)

	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)

	result, err := svc.Query(&dynamodb.QueryInput{
    TableName: aws.String("depploy-users-dev"),
		Limit:     aws.Int64(1),
    KeyConditions: map[string]*dynamodb.Condition{
			"UserId": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("1"),
					},
				},
			},
		},
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

	c.JSON(http.StatusOK, user.Username)
}

func RegisterRoutes() *gin.Engine {

	router := gin.Default()

	user := router.Group("/user")
	{
		user.GET("/:username", username)
	}

	return router
}
