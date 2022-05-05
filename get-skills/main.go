package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"

	"profile/lib"
)

func Handler(_ context.Context) (lib.Response, error) {
	var skills []lib.DbCollection[lib.Skills]
	var val = lib.Skills{}

	return lib.GetCollection(skills, lib.SkillsCollectionName, val)
}

func main() {
	lambda.Start(Handler)
}
