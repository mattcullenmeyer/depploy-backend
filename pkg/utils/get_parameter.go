package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func GetParameter(name string) (string, error) {
	sess, err := session.NewSession()
	if err != nil {
		return "", err
	}
	svc := ssm.New(sess)

	output, err := svc.GetParameter(
		&ssm.GetParameterInput{
			Name: aws.String(name),
			// WithDecryption: aws.Bool(true),
		},
	)
	if err != nil {
		return "", err
	}

	return aws.StringValue(output.Parameter.Value), nil
}
