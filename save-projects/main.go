package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"profile/lib"
	"profile/models"
)

func Handler(_ context.Context, request lib.Request) (lib.Response, error) {
	var projects models.Projects

	lib.PutReqBodyToCollection(lib.ProjectsCollectionName, &projects, &request)

	return lib.ResponseSuccess201()
}

func main() {
	lambda.Start(Handler)
}
