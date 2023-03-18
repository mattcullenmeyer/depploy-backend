package userModel

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
)

type UpdateUserEmailParams struct {
	AccountId string
	Email     string
}

func UpdateUserEmail(args UpdateUserEmailParams) error {
	svc := utils.DynamodbClient()
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	accountIdKey := fmt.Sprintf("ID#%s", strings.ToLower(args.AccountId))
	emailKey := fmt.Sprintf("EMAIL#%s", strings.ToLower(args.Email))

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":gsi1pk": {
				S: aws.String(emailKey),
			},
			":gsi1sk": {
				S: aws.String(emailKey),
			},
			":email": {
				S: aws.String(args.Email),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"PK": {
				S: aws.String(accountIdKey),
			},
			"SK": {
				S: aws.String(accountIdKey),
			},
		},
		UpdateExpression: aws.String("set GSI1PK = :gsi1pk, GSI1SK = :gsi1sk, Email = :email"),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		return err
	}

	return nil
}
