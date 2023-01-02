package userModel

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
)

type GetItemAttributeValues struct {
	Username  string `dynamodbav:"Username"`
	Email     string `dynamodbav:"Email"`
	Password  string `dynamodbav:"Password"`
	CreatedAt string `dynamodbav:"CreatedAt"`
	Verified  bool   `dynamodbav:"Verified"`
}

type FetchUserResult struct {
	Username  string
	Email     string
	Password  string
	CreatedAt string
	Verified  bool
}

func FetchUser(username string) (FetchUserResult, error) {
	svc := utils.DynamodbClient()
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	emptResult := FetchUserResult{}

	key := fmt.Sprintf("ACCOUNT#%s", strings.ToLower(username))

	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"PK": {
				S: aws.String(key),
			},
			"SK": {
				S: aws.String(key),
			},
		},
	}

	getItemOutput, err := svc.GetItem(input)
	if err != nil {
		return emptResult, err
	}

	if getItemOutput.Item == nil {
		return emptResult, nil
	}

	attributeValues := GetItemAttributeValues{}

	err = dynamodbattribute.UnmarshalMap(getItemOutput.Item, &attributeValues)
	if err != nil {
		return emptResult, err
	}

	result := FetchUserResult(attributeValues)

	return result, nil
}
