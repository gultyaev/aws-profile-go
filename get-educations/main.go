package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"profile/models"

	"profile/lib"
)

func Handler(_ context.Context) (lib.Response, error) {
	var educations []models.DbCollection[models.Educations]
	var val = models.Educations{}

	return lib.GetCollection(educations, lib.EducationsCollectionName, val)
}

func main() {
	lambda.Start(Handler)
}
