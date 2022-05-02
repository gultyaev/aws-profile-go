package main

import (
	"bytes"
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
type Request events.APIGatewayProxyRequest

func Handler(_ context.Context, request Request) (Response, error) {
	var item = lib.Skills{}

	log.Printf("Got body %v", request.Body)

	err := json.Unmarshal([]byte(request.Body), &item)

	if err != nil {
		log.Panicf("Error when unmarshalling %v", err)
	}

	log.Printf("Marshalled to %v", item)

	itemCollection := lib.SkillsCollection{
		Collection: lib.SkillsCollectionName,
		Value:      item,
	}

	log.Printf("Prepare item %v", item)

	av, err := dynamodbattribute.MarshalMap(itemCollection)

	if err != nil {
		log.Fatalf("Marshall error mapping value %s", err)
	}

	log.Printf("Marshalled item %v", av)

	input := &dynamodb.PutItemInput{
		TableName: aws.String(lib.DbCollectionsTable),
		Item:      av,
	}

	svc := lib.GetSvc()
	_, err = svc.PutItem(input)

	if err != nil {
		log.Fatalf("Error during put %v", err)
	}

	var buf bytes.Buffer

	body, err := json.Marshal(map[string]interface{}{
		"message": "Save successful!",
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
