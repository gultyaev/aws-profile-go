package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"profile/models"

	"profile/lib"
)

func Handler(_ context.Context) (lib.Response, error) {
	var languages []models.DbCollection[models.Languages]
	var val = models.Languages{}

	return lib.GetCollection(languages, lib.LanguagesCollectionName, val)
}

func main() {
	lambda.Start(Handler)
}
