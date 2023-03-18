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

type FetchUserByEmailParams struct {
	Email string
}

type GetUserByEmailItemAttributeValues struct {
	AccountId  string `dynamodbav:"AccountId"`
	Email      string `dynamodbav:"Email"`
	Password   string `dynamodbav:"Password"`
	CreatedAt  string `dynamodbav:"CreatedAt"`
	Verified   bool   `dynamodbav:"Verified"`
	SuperAdmin bool   `dynamodbav:"SuperAdmin"`
}

type FetchUserByEmailResult struct {
	AccountId  string
	Email      string
	Password   string
	CreatedAt  string
	Verified   bool
	SuperAdmin bool
}

func FetchUserByEmail(args FetchUserByEmailParams) (FetchUserByEmailResult, error) {
	svc := utils.DynamodbClient()
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	emptyResult := FetchUserByEmailResult{}

	emailKey := fmt.Sprintf("EMAIL#%s", strings.ToLower(args.Email))

	pkCondition := expression.Key("GSI1PK").Equal(expression.Value(emailKey))
	skCondition := expression.Key("GSI1SK").Equal(expression.Value(emailKey))
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

	attributeValues := []GetUserByEmailItemAttributeValues{}

	err = dynamodbattribute.UnmarshalListOfMaps(queryOutput.Items, &attributeValues)
	if err != nil {
		return emptyResult, err
	}

	result := FetchUserByEmailResult(attributeValues[0])

	return result, nil
}
