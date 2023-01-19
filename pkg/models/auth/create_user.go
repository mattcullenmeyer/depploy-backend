package authModel

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
)

type CreateUserParams struct {
	Username string
	Email    string
	Password string
}

type UserAttributes struct {
	PK        string `dynamodbav:"PK"`
	SK        string `dynamodbav:"SK"`
	GSI1PK    string `dynamodbav:"GSI1PK"`
	GSI1SK    string `dynamodbav:"GSI1SK"`
	Username  string `dynamodbav:"Username"`
	Email     string `dynamodbav:"Email"`
	Password  string `dynamodbav:"Password"`
	CreatedAt string `dynamodbav:"CreatedAt"`
	Verified  bool   `dynamodbav:"Verified"`
}

func CreateUser(args CreateUserParams) error {
	svc := utils.DynamodbClient()
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	key := fmt.Sprintf("ACCOUNT#%s", strings.ToLower(args.Username))

	user := UserAttributes{
		PK:        key,
		SK:        key,
		GSI1PK:    key, // TODO: this should be user id
		GSI1SK:    key, // TODO: this should be user id
		Username:  args.Username,
		Email:     args.Email,
		Password:  args.Password,
		CreatedAt: time.Now().Format(time.RFC3339), // ISO8601 format for human readability
		Verified:  false,
	}

	item, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return errors.New("error marshalling new user item")
	}

	// Only create new user if username does not already exist
	input := &dynamodb.PutItemInput{
		TableName:           aws.String(tableName),
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(PK)"),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}
