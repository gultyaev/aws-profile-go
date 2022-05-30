package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"profile/models"

	"profile/lib"
)

func Handler(_ context.Context) (lib.Response, error) {
	var skills []models.DbCollection[models.Profile]
	var val = models.Profile{}

	return lib.GetCollection(skills, lib.ProfileCollectionName, val)
}

func main() {
	lambda.Start(Handler)
}
