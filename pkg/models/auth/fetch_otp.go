package authModel

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
)

type FetchOtpParams struct {
	Otp string
}

type QueryAttributeValues struct {
	AccountId  string `dynamodbav:"AccountId"`
	Password   string `dynamodbav:"Password"`
	Email      string `dynamodbav:"Email"`
	Expiration string `dynamodbav:"Expiration"`
}

type FetchOtpResult struct {
	AccountId  string
	Password   string
	Email      string
	Expiration string
}

func FetchOtp(args FetchOtpParams) (FetchOtpResult, error) {
	svc := utils.DynamodbClient()
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	emptyResult := FetchOtpResult{}

	otpKey := fmt.Sprintf("OTP#%s", strings.ToLower(args.Otp))

	pkCondition := expression.Key("PK").Equal(expression.Value(otpKey))
	skCondition := expression.Key("SK").Equal(expression.Value(otpKey))
	keyCondition := pkCondition.And(skCondition)

	// Only include verification codes where expiration is in the future
	now := time.Now().Unix()
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

	result := FetchOtpResult(attributeValues[0])

	return result, nil
}
