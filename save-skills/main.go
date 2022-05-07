package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"profile/models"

	"profile/lib"
)

func Handler(_ context.Context, request lib.Request) (lib.Response, error) {
	var item *models.Skills

	lib.PutReqBodyToCollection(lib.SkillsCollectionName, &item, &request)

	return lib.ResponseSuccess201()
}

func main() {
	lambda.Start(Handler)
}
