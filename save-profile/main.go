package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"profile/lib"
)

func Handler(_ context.Context, request lib.Request) (lib.Response, error) {
	var profile lib.Profile

	lib.PutReqBodyToCollection(lib.ProfileCollectionName, &profile, &request)

	return lib.ResponseSuccess201()
}

func main() {
	lambda.Start(Handler)
}
