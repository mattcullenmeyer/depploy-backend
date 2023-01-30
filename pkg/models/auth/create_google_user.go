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

type CreateGoogleUserParams struct {
	AccountId string
	Email     string
	Verified  bool
	Name      string
}

type GoogleUserAttributes struct {
	PK                 string `dynamodbav:"PK"`
	SK                 string `dynamodbav:"SK"`
	GSI1PK             string `dynamodbav:"GSI1PK"`
	GSI1SK             string `dynamodbav:"GSI1SK"`
	Username           string `dynamodbav:"Username"`
	AccountId          string `dynamodbav:"AccountId"`
	Email              string `dynamodbav:"Email"`
	CreatedAt          string `dynamodbav:"CreatedAt"`
	Verified           bool   `dynamodbav:"Verified"`
	Superuser          bool   `dynamodbav:"Superuser"`
	Type               string `dynamodbav:"Type"`
	RegistrationMethod string `dynamodbav:"RegistrationMethod"`
}

func CreateGoogleUser(args CreateGoogleUserParams) error {
	svc := utils.DynamodbClient()
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	key := fmt.Sprintf("ACCOUNT#%s", strings.ToLower(args.AccountId))
	gsi1Key := fmt.Sprintf("ID#%s", args.AccountId)

	user := GoogleUserAttributes{
		PK:                 key,
		SK:                 key,
		GSI1PK:             gsi1Key,
		GSI1SK:             gsi1Key,
		Username:           "",
		AccountId:          args.AccountId,
		Email:              args.Email,
		CreatedAt:          time.Now().Format(time.RFC3339), // ISO8601 format for human readability
		Verified:           args.Verified,
		Superuser:          false,
		Type:               "User Account",
		RegistrationMethod: "Google",
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
