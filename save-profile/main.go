package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"profile/lib"
	"profile/models"
)

func Handler(_ context.Context, request lib.Request) (lib.Response, error) {
	var profile models.Profile

	lib.PutReqBodyToCollection(lib.ProfileCollectionName, &profile, &request)

	return lib.ResponseSuccess201()
}

func main() {
	lambda.Start(Handler)
}
