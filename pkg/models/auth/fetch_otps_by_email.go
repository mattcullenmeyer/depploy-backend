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

type FetchOtpsByEmailParams struct {
	Email string
}

type FetchOtpsByEmailResult struct {
	AccountId  string
	Password   string
	Email      string
	Expiration string
}

func FetchOtpsByEmail(args FetchOtpsByEmailParams) ([]FetchOtpsByEmailResult, error) {
	svc := utils.DynamodbClient()
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	emptyResult := []FetchOtpsByEmailResult{}

	emailKey := fmt.Sprintf("OTP#%s", strings.ToLower(args.Email))

	pkCondition := expression.Key("GSI2PK").Equal(expression.Value(emailKey))
	skCondition := expression.Key("GSI2SK").Equal(expression.Value(emailKey))
	keyCondition := pkCondition.And(skCondition)

	// Only include verification codes where expiration is in the future
	now := time.Now().Unix()
	filter := expression.Name("TTL").GreaterThan(expression.Value(now))

	expr, err := expression.NewBuilder().
		WithKeyCondition(keyCondition).
		WithFilter(filter).
		Build()
	if err != nil {
		return emptyResult, err
	}

	input := &dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		IndexName:                 aws.String("GSI2"),
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

	result := []FetchOtpsByEmailResult{}

	for _, i := range queryOutput.Items {
		item := FetchOtpsByEmailResult{}

		err = dynamodbattribute.UnmarshalMap(i, &item)
		if err != nil {
			return emptyResult, err
		}

		result = append(result, item)
	}

	return result, nil
}
