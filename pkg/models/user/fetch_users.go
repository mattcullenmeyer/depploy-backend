package userModel

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
)

type UserResult struct {
	AccountId string
	Username  string
	Email     string
	Verified  bool
}

type FetchUsersResult struct {
	Users []UserResult
	Next  string
}

func FetchUsers(limit int64, key string) (FetchUsersResult, error) {
	svc := utils.DynamodbClient()
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	emptyResult := FetchUsersResult{}

	exclusiveStartKey, err := utils.DecodeLastEvaluatedKey(key)
	if err != nil {
		return emptyResult, err
	}

	projection := expression.NamesList(
		expression.Name("AccountId"),
		expression.Name("Username"),
		expression.Name("Email"),
		expression.Name("Verified"),
	)

	expr, err := expression.NewBuilder().
		WithProjection(projection).
		Build()
	if err != nil {
		return emptyResult, err
	}

	input := &dynamodb.ScanInput{
		TableName:                 aws.String(tableName),
		IndexName:                 aws.String("GSI1"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ProjectionExpression:      expr.Projection(),
		Limit:                     aws.Int64(limit),
	}

	if key != "" {
		input.ExclusiveStartKey = exclusiveStartKey
	}

	output, err := svc.Scan(input)
	if err != nil {
		return emptyResult, err
	}

	lastEvaluatedKey, err := utils.EncodeLastEvaluatedKey(output.LastEvaluatedKey)
	if err != nil {
		return emptyResult, err
	}

	userResult := []UserResult{}

	for _, i := range output.Items {
		item := UserResult{}

		err = dynamodbattribute.UnmarshalMap(i, &item)
		if err != nil {
			return emptyResult, err
		}

		userResult = append(userResult, item)
	}

	result := FetchUsersResult{
		Users: userResult,
		Next:  lastEvaluatedKey,
	}

	return result, nil
}
