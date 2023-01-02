package authModel

import (
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

	emptyResult := FetchVerificationCodeResult{}

	now := time.Now().Unix()

	pkCondition := expression.Key("PK").Equal(expression.Value(otp))
	skCondition := expression.Key("SK").Equal(expression.Value(otp))
	keyCondition := pkCondition.And(skCondition)

	// Only include verification codes where expiration is in the future
	filter := expression.Name("TTL").GreaterThan(expression.Value(now))

	// Combine the key condition and filter together as a DynamoDB expression builder
	expr, err := expression.NewBuilder().
		WithKeyCondition(keyCondition).
		WithFilter(filter).
		Build()
	if err != nil {
		return emptyResult, err
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
		return emptyResult, err
	}

	if len(queryOutput.Items) == 0 {
		return emptyResult, nil
	}

	attributeValues := []QueryAttributeValues{}

	err = dynamodbattribute.UnmarshalListOfMaps(queryOutput.Items, &attributeValues)
	if err != nil {
		return emptyResult, err
	}

	result := FetchVerificationCodeResult(attributeValues[0])

	return result, nil
}
