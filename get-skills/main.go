package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"profile/models"

	"profile/lib"
)

func Handler(_ context.Context) (lib.Response, error) {
	var skills []models.DbCollection[models.Skills]
	var val = models.Skills{}

	return lib.GetCollection(skills, lib.SkillsCollectionName, val)
}

func main() {
	lambda.Start(Handler)
}
