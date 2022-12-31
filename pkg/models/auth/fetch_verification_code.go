package authModel

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
)

type QueryAttributeValues struct {
	Username string `dynamodbav:"Username"`
	Email    string `dynamodbav:"Email"`
}

type FetchVerificationCodeResult struct {
	Username string
	Email    string
}

func FetchVerificationCode(otp string) (FetchVerificationCodeResult, error) {
	svc := utils.DynamodbClient()
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	// Return empty result if error
	emptyResult := FetchVerificationCodeResult{}

	ttl := time.Now().Unix()

	keyCondition := expression.Key("PK").Equal(expression.Value(otp))

	// Exclude expired verification codes
	filter := expression.Name("TTL").GreaterThan(expression.Value(ttl))

	// Combine the key condition and filter together as a DynamoDB expression builder
	expr, err := expression.NewBuilder().
		WithKeyCondition(keyCondition).
		WithFilter(filter).
		Build()
	if err != nil {
		log.Println(err)
		return emptyResult, errors.New("something went wrong")
	}

	input := &dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}

	queryOutput, err := svc.Query(input)
	if err != nil {
		log.Fatalf("Got error calling QueryInput: %s", err)
	}

	if len(queryOutput.Items) == 0 {
		return emptyResult, errors.New("invalid verification code")
	}

	attributeValues := []QueryAttributeValues{}

	err = dynamodbattribute.UnmarshalListOfMaps(queryOutput.Items, &attributeValues)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	result := FetchVerificationCodeResult{
		Username: attributeValues[0].Username,
		Email:    attributeValues[0].Email,
	}

	return result, nil
}
