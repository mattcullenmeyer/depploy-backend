package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"fmt"
	"log"
)

func CreateTable() {
	svc := CreateClient()

	// Create table
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("PK"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("SK"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("GSI1PK"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("GSI1SK"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("GSI2PK"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("GSI2SK"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("PK"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("SK"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String("Depploy"),
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String("GSI1"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("GSI1PK"),
						KeyType:       aws.String("HASH"),
					},
					{
						AttributeName: aws.String("GSI1SK"),
						KeyType:       aws.String("RANGE"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(10),
					WriteCapacityUnits: aws.Int64(10),
				},
			},
			{
				IndexName: aws.String("GSI2"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("GSI2PK"),
						KeyType:       aws.String("HASH"),
					},
					{
						AttributeName: aws.String("GSI2SK"),
						KeyType:       aws.String("RANGE"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(10),
					WriteCapacityUnits: aws.Int64(10),
				},
			},
		},
	}

	_, err := svc.CreateTable(input)
	if err != nil {
		log.Fatalf("Got error calling CreateTable: %s", err)
	}

	fmt.Println("Created the table")
}
