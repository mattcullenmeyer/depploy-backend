package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func DeleteTable() {
	svc := CreateClient()

	// Delete table
	input := &dynamodb.DeleteTableInput{
		TableName: aws.String("Depploy"),
	}

	_, err := svc.DeleteTable(input)

	if err != nil {
		log.Fatalf("Got error calling DeleteTable: %s", err)
	}

	fmt.Println("Deleted the existing table")
}
