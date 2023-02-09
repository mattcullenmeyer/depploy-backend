package userModel

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
)

type FetchUserByUsernameParams struct {
	Username string
}

type GetUserByUsernameItemAttributeValues struct {
	Username   string `dynamodbav:"Username"`
	AccountId  string `dynamodbav:"AccountId"`
	Email      string `dynamodbav:"Email"`
	Password   string `dynamodbav:"Password"`
	CreatedAt  string `dynamodbav:"CreatedAt"`
	Verified   bool   `dynamodbav:"Verified"`
	SuperAdmin bool   `dynamodbav:"SuperAdmin"`
}

type FetchUserByUsernameResult struct {
	Username   string
	AccountId  string
	Email      string
	Password   string
	CreatedAt  string
	Verified   bool
	SuperAdmin bool
}

func FetchUserByUsername(args FetchUserByUsernameParams) (FetchUserByUsernameResult, error) {
	svc := utils.DynamodbClient()
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	emptyResult := FetchUserByUsernameResult{}

	accountNameKey := fmt.Sprintf("ACCOUNT#%s", strings.ToLower(args.Username))

	pkCondition := expression.Key("GSI1PK").Equal(expression.Value(accountNameKey))
	skCondition := expression.Key("GSI1SK").Equal(expression.Value(accountNameKey))
	keyCondition := pkCondition.And(skCondition)

	expr, err := expression.NewBuilder().
		WithKeyCondition(keyCondition).
		Build()
	if err != nil {
		return emptyResult, err
	}

	input := &dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		IndexName:                 aws.String("GSI1"),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}

	queryOutput, err := svc.Query(input)
	if err != nil {
		return emptyResult, err
	}

	if len(queryOutput.Items) == 0 {
		return emptyResult, nil
	}

	attributeValues := []GetUserByUsernameItemAttributeValues{}

	err = dynamodbattribute.UnmarshalListOfMaps(queryOutput.Items, &attributeValues)
	if err != nil {
		return emptyResult, err
	}

	result := FetchUserByUsernameResult(attributeValues[0])

	return result, nil
}
