package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

// aws sts

var credentials = map[string]string{
	// Mongo Paths + URI
	"MONGODB_URI":     "mongodb://localhost:27017",
	"SHARED_LIB_PATH": "<path to crypt_shared library>",
	// AWS Credentials
	"AWS_ACCESS_KEY_ID":     "<your AWS access key ID here>",
	"AWS_SECRET_ACCESS_KEY": "<your AWS secret access key here>",
	"AWS_KEY_REGION":        "<your AWS key region>",
	"AWS_KEY_ARN":           "<your AWS key ARN>",
}

// return credentials object and ensure it has been populated
func GetCredentials() map[string]string {
	stscli := sts.New(session.New())
}
