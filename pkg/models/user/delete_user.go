package userModel

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
)

type DeleteUserParams struct {
	AccountId string
}

// WARNING: Do NOT add this to an API endpoint for security reasons
// This is only used through CLI script
func DeleteUser(args DeleteUserParams) error {
	svc := utils.DynamodbClient()
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	accountIdKey := fmt.Sprintf("ID#%s", strings.ToLower(args.AccountId))

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"PK": {
				S: aws.String(accountIdKey),
			},
			"SK": {
				S: aws.String(accountIdKey),
			},
		},
		TableName: aws.String(tableName),
	}

	_, err := svc.DeleteItem(input)
	if err != nil {
		return err
	}

	return nil
}
