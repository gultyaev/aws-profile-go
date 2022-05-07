package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"profile/models"

	"profile/lib"
)

func Handler(_ context.Context) (lib.Response, error) {
	var projects []models.DbCollection[models.Projects]
	var val = models.Projects{}

	return lib.GetCollection(projects, lib.ProjectsCollectionName, val)
}

func main() {
	lambda.Start(Handler)
}
