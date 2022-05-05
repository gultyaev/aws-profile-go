package lib

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
)

type Request events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse
type DynamoMap map[string]*dynamodb.AttributeValue

func UnmarshalRequestOrFail(req *Request, tgt any) {
	log.Printf("Got body %v", &req.Body)

	err := json.Unmarshal([]byte(req.Body), &tgt)

	if err != nil {
		log.Panicf("Error when unmarshalling %v", err)
	}

	log.Printf("Unarshalled to %v", &tgt)
}

func ToCollection[T any](collectionName string, value T) DbCollection[T] {
	return DbCollection[T]{
		Collection: collectionName,
		Value:      value,
	}
}

func MarshalDynamoDB[T any](value T) (result DynamoMap) {
	result, err := dynamodbattribute.MarshalMap(value)

	if err != nil {
		log.Fatalf("Marshall error mapping value %s", err)
	}

	log.Printf("Marshalled item %v", result)

	return
}

func PutCollection(table string, collection DynamoMap) {
	input := &dynamodb.PutItemInput{
		TableName: aws.String(table),
		Item:      collection,
	}

	svc := GetSvc()
	_, err := svc.PutItem(input)

	if err != nil {
		log.Fatalf("Error during put %v", err)
	}
}

func PutReqBodyToCollection[T any](collection string, value *T, request *Request) {
	UnmarshalRequestOrFail(request, &value)
	skillsCollection := ToCollection(collection, &value)
	av := MarshalDynamoDB(skillsCollection)
	PutCollection(DbCollectionsTable, av)
}

func ResponseSuccess201() (Response, error) {
	var buf bytes.Buffer

	body, err := json.Marshal(map[string]interface{}{
		"message": "Created",
	})

	if err != nil {
		log.Panicf("Error during marshalling response %v", err)
	}

	json.HTMLEscape(&buf, body)

	return Response{
		StatusCode:      201,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func GetCollection[T any](collection []DbCollection[T], collectionName string, value T) (Response, error) {
	svc := GetSvc()

	input := &dynamodb.ScanInput{
		TableName:        aws.String(DbCollectionsTable),
		FilterExpression: aws.String("#name = :skills"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":skills": {S: aws.String(collectionName)},
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

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &collection)

	if err != nil {
		log.Panicf("Error during unmarshalling %v", err)
	}

	if len(collection) != 0 {
		value = collection[0].Value
	}

	res, err := json.Marshal(value)

	if err != nil {
		log.Panicf("Error during marshalling result %v", res)
	}

	return Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(res),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
