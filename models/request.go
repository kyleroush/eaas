package models

import "github.com/aws/aws-lambda-go/events"

//Input are the possible inputs that lambda takes
type Input struct {
	Format      string
	ContentType string
	Request     events.APIGatewayProxyRequest
	Inputs      map[string]string
}
