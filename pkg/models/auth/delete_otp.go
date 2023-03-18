package authModel

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
)

type DeleteOtpParams struct {
	Otp string
}

type DeleteOtpResult struct {
	AccountId  string
	Password   string
	Email      string
	Expiration string
}

func DeleteOtp(args DeleteOtpParams) (DeleteOtpResult, error) {
	svc := utils.DynamodbClient()
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	emptyResult := DeleteOtpResult{}

	now := time.Now().Unix()
	nowString := strconv.FormatInt(now, 10)

	otpKey := fmt.Sprintf("OTP#%s", strings.ToLower(args.Otp))

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"PK": {
				S: aws.String(otpKey),
			},
			"SK": {
				S: aws.String(otpKey),
			},
		},
		ConditionExpression: aws.String("#ttl > :now"),
		ExpressionAttributeNames: map[string]*string{
			"#ttl": aws.String("TTL"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":now": {
				N: aws.String(nowString),
			},
		},
		ReturnValues: aws.String("ALL_OLD"),
	}

	deleteItemOutput, err := svc.DeleteItem(input)
	if err != nil {
		return emptyResult, err
	}

	deleteOtpResult := DeleteOtpResult{}
	err = dynamodbattribute.UnmarshalMap(deleteItemOutput.Attributes, &deleteOtpResult)
	if err != nil {
		return emptyResult, err
	}

	return deleteOtpResult, nil
}
