package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/service/dynamodb"

	"profile/lib"
)

type Response events.APIGatewayProxyResponse

func Handler(_ context.Context) (Response, error) {
	log.Println("Get get-skills request")

	svc := lib.GetSvc()
	input := &dynamodb.ScanInput{
		TableName:        aws.String(lib.DbCollectionsTable),
		FilterExpression: aws.String("#name = :skills"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":skills": {S: aws.String(lib.SkillsCollectionName)},
		},
		ExpressionAttributeNames: map[string]*string{
			"#name":  aws.String("Collection"),
			"#value": aws.String("Value"),
		},
		ConsistentRead:       aws.Bool(true),
		ProjectionExpression: aws.String("#value"),
	}

	log.Println("Start scanning collections")

	result, err := svc.Scan(input)

	if err != nil {
		log.Panicf("Error during scan operation %v", err)
	}

	log.Printf("Scan result %v", result.Items)

	var skills []lib.SkillsCollection

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &skills)

	if err != nil {
		log.Panicf("Error during unmarshalling %v", err)
	}

	var val lib.Skills

	if len(skills) == 0 {
		val = lib.Skills{}
	} else {
		val = skills[0].Value
	}

	res, err := json.Marshal(val)

	if err != nil {
		log.Panicf("Error during marshalling result %v", res)
	}

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(res),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
