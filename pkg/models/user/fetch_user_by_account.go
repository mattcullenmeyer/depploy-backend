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

type FetchUserByAccountParams struct {
	AccountId string
}

type GetUserByAccountItemAttributeValues struct {
	AccountId  string `dynamodbav:"AccountId"`
	Email      string `dynamodbav:"Email"`
	Password   string `dynamodbav:"Password"`
	CreatedAt  string `dynamodbav:"CreatedAt"`
	Verified   bool   `dynamodbav:"Verified"`
	SuperAdmin bool   `dynamodbav:"SuperAdmin"`
}

type FetchUserByAccountResult struct {
	AccountId  string
	Email      string
	Password   string
	CreatedAt  string
	Verified   bool
	SuperAdmin bool
}

func FetchUserByAccount(args FetchUserByAccountParams) (FetchUserByAccountResult, error) {
	svc := utils.DynamodbClient()
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	emptyResult := FetchUserByAccountResult{}

	accountIdKey := fmt.Sprintf("ID#%s", strings.ToLower(args.AccountId))

	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"PK": {
				S: aws.String(accountIdKey),
			},
			"SK": {
				S: aws.String(accountIdKey),
			},
		},
	}

	getItemOutput, err := svc.GetItem(input)
	if err != nil {
		return emptyResult, err
	}

	if getItemOutput.Item == nil {
		return emptyResult, nil
	}

	attributeValues := GetUserByAccountItemAttributeValues{}

	err = dynamodbattribute.UnmarshalMap(getItemOutput.Item, &attributeValues)
	if err != nil {
		return emptyResult, err
	}

	result := FetchUserByAccountResult(attributeValues)

	return result, nil
}
