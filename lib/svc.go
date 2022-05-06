package lib

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

var savedDynamoSession *dynamodb.DynamoDB
var savedS3Session *s3.S3

func GetDynamoDBSession() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	if savedDynamoSession == nil {
		savedDynamoSession = dynamodb.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))
	}

	log.Printf("Created session")

	return savedDynamoSession
}

func GetS3Session() *s3.S3 {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	if savedS3Session == nil {
		savedS3Session = s3.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))
	}

	log.Printf("Created session")

	return savedS3Session
}
