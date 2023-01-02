package authModel

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
)

func UpdateUserVerified(username string) error {
	svc := utils.DynamodbClient()
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	key := fmt.Sprintf("ACCOUNT#%s", strings.ToLower(username))

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":verified": {
				BOOL: aws.Bool(true),
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"PK": {
				S: aws.String(key),
			},
			"SK": {
				S: aws.String(key),
			},
		},
		UpdateExpression: aws.String("set Verified = :verified"),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		return err
	}

	return nil
}
