package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"net/url"
	"profile/lib"
	"time"
)

func Handler(_ context.Context, request lib.Request) (lib.Response, error) {
	svc := lib.GetS3Session()

	imageName := request.QueryStringParameters["image"]
	imageNameUnescaped, err := url.QueryUnescape(imageName)

	if err != nil {
		log.Panicf("Cannot parse query %v", err)
	}

	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(lib.ImagesBucket),
		Key:    aws.String(imageNameUnescaped),
	})

	imgUrl, err := req.Presign(15 * time.Minute)

	if err != nil {
		log.Panicf("Error during Presign %v", err)
	}

	res, _ := json.Marshal(lib.PresignImage{
		Url:       lib.GetBucketUrl() + "/" + imageName,
		UploadUrl: imgUrl,
	})

	return lib.Response{
		StatusCode:      200,
		Body:            string(res),
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
