package userModel

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
)

type UpdateUserSuperuserParams struct {
	AccountId string
	Access    bool
}

// WARNING: Do NOT add this to an API endpoint for security reasons
// This is only used through CLI script
func UpdateUserSuperuser(args UpdateUserSuperuserParams) error {
	svc := utils.DynamodbClient()
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	accountIdKey := fmt.Sprintf("ID#%s", strings.ToLower(args.AccountId))

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":superuser": {
				BOOL: aws.Bool(args.Access),
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
		UpdateExpression: aws.String("set Superuser = :superuser"),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		return err
	}

	return nil
}
