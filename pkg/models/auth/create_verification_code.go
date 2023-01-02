package authModel

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
)

type CreateVerificationCodeParams struct {
	Otp      string
	Username string
	Email    string
}

type VerificationCodeAttributes struct {
	PK         string `dynamodbav:"PK"`
	SK         string `dynamodbav:"SK"`
	Username   string `dynamodbav:"Username"`
	Email      string `dynamodbav:"Email"`
	Expiration string `dynamodbav:"Expiration"`
	TTL        int64  `dynamodbav:"TTL"`
}

func CreateVerificationCode(args CreateVerificationCodeParams) error {
	svc := utils.DynamodbClient()
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	ttl := time.Now().Add(time.Hour * 24)

	verificationCode := VerificationCodeAttributes{
		PK:         args.Otp,
		SK:         args.Otp,
		Username:   strings.ToLower(args.Username),
		Email:      args.Email,
		Expiration: ttl.Format(time.RFC3339), // ISO8601 format for human readability
		TTL:        ttl.Unix(),
	}

	item, err := dynamodbattribute.MarshalMap(verificationCode)
	if err != nil {
		return errors.New("error marshalling verification code")
	}

	input := &dynamodb.PutItemInput{
		TableName:           aws.String(tableName),
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(PK)"),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}
