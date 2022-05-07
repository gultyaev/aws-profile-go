package main

import (
	"context"
	"encoding/base64"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"profile/lib"
	"strings"
)

func generatePolicy(principalId, effect, resource string) events.APIGatewayCustomAuthorizerResponse {
	authResponse := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalId}

	if effect != "" && resource != "" {
		authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resource},
				},
			},
		}
	}

	return authResponse
}

func handleRequest(ctx context.Context, event events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	if event.Type != "TOKEN" {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	token := event.AuthorizationToken

	if token == "" {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	tokenValues := strings.Split(token, " ")

	if len(tokenValues) != 2 || tokenValues[0] != "Basic" {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Error: Invalid token")
	}

	parsedToken, err := base64.StdEncoding.DecodeString(tokenValues[1])

	if err != nil {
		log.Printf("Error parsing token %v", err)
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	creds := strings.Split(string(parsedToken), ":")

	if len(creds) != 2 {
		return generatePolicy("user", "Deny", event.MethodArn), nil
	} else if creds[0] != lib.Username || creds[1] != lib.Password {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	return generatePolicy(creds[0], "Allow", event.MethodArn), nil
}

func main() {
	lambda.Start(handleRequest)
}
