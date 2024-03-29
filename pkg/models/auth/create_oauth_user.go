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

type CreateOauthUserParams struct {
	AccountId          string
	Email              string
	Name               string
	RegistrationMethod string
}

type OauthUserAttributes struct {
	PK                 string `dynamodbav:"PK"`
	SK                 string `dynamodbav:"SK"`
	GSI1PK             string `dynamodbav:"GSI1PK"`
	GSI1SK             string `dynamodbav:"GSI1SK"`
	Username           string `dynamodbav:"Username"`
	AccountId          string `dynamodbav:"AccountId"`
	Email              string `dynamodbav:"Email"`
	CreatedAt          string `dynamodbav:"CreatedAt"`
	Verified           bool   `dynamodbav:"Verified"`
	SuperAdmin         bool   `dynamodbav:"SuperAdmin"`
	Type               string `dynamodbav:"Type"`
	RegistrationMethod string `dynamodbav:"RegistrationMethod"`
}

func CreateOauthUser(args CreateOauthUserParams) error {
	svc := utils.DynamodbClient()
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	accountIdKey := fmt.Sprintf("ID#%s", strings.ToLower(args.AccountId))
	accountNameKey := fmt.Sprintf("ACCOUNT#%s", strings.ToLower(args.AccountId))

	user := OauthUserAttributes{
		PK:                 accountIdKey,
		SK:                 accountIdKey,
		GSI1PK:             accountNameKey,
		GSI1SK:             accountNameKey,
		Username:           "",
		AccountId:          args.AccountId,
		Email:              args.Email,
		CreatedAt:          time.Now().Format(time.RFC3339), // ISO8601 format for human readability
		Verified:           true,
		SuperAdmin:         false,
		Type:               "User Account",
		RegistrationMethod: args.RegistrationMethod,
	}

	item, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return errors.New("error marshalling new user item")
	}

	// Only create new user if username does not already exist
	input := &dynamodb.PutItemInput{
		TableName:           aws.String(tableName),
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(PK) AND attribute_not_exists(GSI1PK)"),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}
