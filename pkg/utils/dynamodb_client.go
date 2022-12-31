package utils

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func DynamodbClient() *dynamodb.DynamoDB {
	endpoint := os.Getenv("DYNAMODB_ENDPOINT")

	config := &aws.Config{
		Endpoint: aws.String(endpoint), // uses default generated endpoint if an empty string
	}

	sess := session.Must(session.NewSession(config))

	return dynamodb.New(sess)
}
