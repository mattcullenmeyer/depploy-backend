package userModel

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
)

type UpdateAccountUsernameParams struct {
	Username  string
	AccountId string
}

func UpdateAccountUsername(args UpdateAccountUsernameParams) error {
	svc := utils.DynamodbClient()
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	accountIdKey := fmt.Sprintf("ID#%s", strings.ToLower(args.AccountId))
	accountNameKey := fmt.Sprintf("ACCOUNT#%s", strings.ToLower(args.Username))

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":gsi1pk": {
				S: aws.String(accountNameKey),
			},
			":gsi1sk": {
				S: aws.String(accountNameKey),
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
		UpdateExpression: aws.String("set GSI1PK = :gsi1pk, GSI1SK = :gsi1sk"),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		return err
	}

	return nil
}
