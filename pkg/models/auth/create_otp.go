package authModel

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
)

type CreateOtpParams struct {
	Otp       string
	AccountId string
	Email     string
}

type OtpAttributes struct {
	PK         string `dynamodbav:"PK"`
	SK         string `dynamodbav:"SK"`
	GSI2PK     string `dynamodbav:"GSI2PK"`
	GSI2SK     string `dynamodbav:"GSI2SK"`
	AccountId  string `dynamodbav:"AccountId"`
	Password   string `dynamodbav:"Password"`
	Email      string `dynamodbav:"Email"`
	Expiration string `dynamodbav:"Expiration"`
	TTL        int64  `dynamodbav:"TTL"`
}

func CreateOtp(args CreateOtpParams) error {
	svc := utils.DynamodbClient()
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	otpKey := fmt.Sprintf("OTP#%s", strings.ToLower(args.Otp))
	emailKey := fmt.Sprintf("OTP#%s", strings.ToLower(args.Email))
	ttl := time.Now().Add(time.Minute * 15)

	otp := OtpAttributes{
		PK:         otpKey,
		SK:         otpKey,
		GSI2PK:     emailKey,
		GSI2SK:     emailKey,
		AccountId:  args.AccountId,
		Password:   args.Otp,
		Email:      args.Email,
		Expiration: ttl.Format(time.RFC3339), // ISO8601 format for human readability
		TTL:        ttl.Unix(),
	}

	item, err := dynamodbattribute.MarshalMap(otp)
	if err != nil {
		return errors.New("error marshalling verification code")
	}

	input := &dynamodb.PutItemInput{
		TableName:           aws.String(tableName),
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(PK) AND attribute_not_exists(GSI2PK)"),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}
