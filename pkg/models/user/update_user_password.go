package userModel

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
)

type UpdateUserPasswordParams struct {
	AccountId string
	Password  string
}

func UpdateUserPassword(args UpdateUserPasswordParams) error {
	svc := utils.DynamodbClient()
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	accountIdKey := fmt.Sprintf("ID#%s", strings.ToLower(args.AccountId))

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":password": {
				S: aws.String(args.Password),
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"PK": {
				S: aws.String(accountIdKey),
			},
			"SK": {
				S: aws.String(accountIdKey),
			},
		},
		UpdateExpression: aws.String("set Password = :password"),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		return err
	}

	return nil
}
