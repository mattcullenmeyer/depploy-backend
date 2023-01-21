package userModel

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
)

// Update the dynamodb projection below if you edit the UserResult type struct here
type UserResult struct {
	Username  string
	AccountId string
	Email     string
	Verified  bool
	Superuser bool
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

	// update the UserResult struct type above if you edit the projection here
	projection := expression.NamesList(
		expression.Name("Username"),
		expression.Name("AccountId"),
		expression.Name("Email"),
		expression.Name("Verified"),
		expression.Name("Superuser"),
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
