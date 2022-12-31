package authModel

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
)

type GetItemAttributeValues struct {
	Password string `dynamodbav:"Password"`
	Email    string `dynamodbav:"Email"`
	Verified bool   `dynamodbav:"Verified"`
}

type FetchUserPasswordResult struct {
	Password string
	Verified bool
}

func FetchUserPassword(username string) FetchUserPasswordResult {
	svc := utils.DynamodbClient()
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	key := fmt.Sprintf("ACCOUNT#%s", username)

	item := &dynamodb.GetItemInput{
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

	getItemOutput, err := svc.GetItem(item)
	if err != nil {
		msg := "Got error calling GetItem:" + err.Error()
		log.Println(msg)
	}

	if getItemOutput.Item == nil {
		msg := "Could not find '" + username + "'"
		log.Println(msg)
	}

	attributeValues := GetItemAttributeValues{}

	err = dynamodbattribute.UnmarshalMap(getItemOutput.Item, &attributeValues)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	result := FetchUserPasswordResult{
		Password: attributeValues.Password,
		Verified: attributeValues.Verified,
	}

	return result
}
