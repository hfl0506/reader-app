package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type AccessPayload struct {
	AwsRegion          string
	AwsAccessKeyId     string
	AwsSecretAccessKey string
}

func ConnectAws(payload *AccessPayload) *session.Session {
	sess, err := session.NewSession(
		&aws.Config{
			Region: &payload.AwsRegion,
			Credentials: credentials.NewStaticCredentials(
				payload.AwsAccessKeyId,
				payload.AwsSecretAccessKey,
				"",
			),
		},
	)
	if err != nil {
		panic(err)
	}

	return sess
}
