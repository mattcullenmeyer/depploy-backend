package utils

import (
	"encoding/base64"
	"encoding/json"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// https://stackoverflow.com/questions/68308139/how-to-serialize-lastevaluatedkey-from-dynamodbs-golang-sdk

func EncodeLastEvaluatedKey(input map[string]*dynamodb.AttributeValue) (string, error) {
	if len(input) == 0 {
		return "", nil
	}

	inputMap := map[string]string{}
	if err := dynamodbattribute.UnmarshalMap(input, &inputMap); err != nil {
		return "", err
	}

	bytesJSON, err := json.Marshal(inputMap)
	if err != nil {
		return "", err
	}

	output := base64.StdEncoding.EncodeToString(bytesJSON)

	return output, nil
}

func DecodeLastEvaluatedKey(input string) (map[string]*dynamodb.AttributeValue, error) {
	if input == "" {
		return nil, nil
	}

	bytesJSON, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return nil, err
	}

	outputJSON := map[string]string{}
	if err := json.Unmarshal(bytesJSON, &outputJSON); err != nil {
		return nil, err
	}

	output, err := dynamodbattribute.MarshalMap(outputJSON)
	if err != nil {
		return nil, err
	}

	return output, nil
}
