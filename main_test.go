package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	request := events.APIGatewayProxyRequest{
		Headers: map[string]string{
			"accepts": "application/json",
		},
		Body: "token=srVPRsp4HXUUjUygIgIjzijA&team_id=T5SM23F44&team_domain=roushhouse&channel_id=D5TDW8G22&channel_name=directmessage&user_id=U5SNE3VGV&user_name=roush&command=%2Fexcuseme&text=dog%2Cali&response_url=https%3A%2F%2Fhooks.slack.com%2Fcommands%2FT5SM23F44%2F548585650213%2FJbblTwNx4vmrqsP8k2XU7ex1&trigger_id=547775724865.196716117140.d19dd8ca00a61df78e29a0b906de53e0",
		QueryStringParameters: map[string]string{
			"to":     "Scott",
			"from":   "kyle",
			"excuse": "dog",
			// "format": "slack",
		}}
	expectedResponse := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Access-Control-Allow-Origin": "*",
			"Access-Control-Allow-Method": "*",
			"Access-Control-Allow-Header": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
		},
		Body: "kyle",
	}

	response, err := Handler(request)

	assert.Equal(t, response.Headers, expectedResponse.Headers)
	assert.Contains(t, response.Body, expectedResponse.Body)
	assert.Equal(t, err, nil)

}

func TestSlack(t *testing.T) {

	request := events.APIGatewayProxyRequest{
		Body: "token=srVPRsp4HXUUjUygIgIjzijA&" +
			"team_id=T5SM23F44&" +
			"team_domain=roushhouse&" +
			"channel_id=D5TDW8G22&" +
			"channel_name=directmessage&" +
			"user_id=U5SNE3VGV&" +
			"user_name=roush&command=%2Fexcuseme&" +
			"text=dog%2Cali&" +
			"response_url=https%3A%2F%2Fhooks.slack.com%2Fcommands%2FT5SM23F44%2F548585650213%2FJbblTwNx4vmrqsP8k2XU7ex1&" +
			"trigger_id=547775724865.196716117140.d19dd8ca00a61df78e29a0b906de53e0",
		QueryStringParameters: map[string]string{
			"format": "slack",
		}}
	expectedResponse := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Access-Control-Allow-Origin": "*",
			"Access-Control-Allow-Method": "*",
			"Access-Control-Allow-Header": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
		},
		Body: "{\"text\":\"Dear ali, My dog ate my homework. Sincerly roush\"}",
	}

	response, err := Handler(request)

	assert.Equal(t, response.Headers, expectedResponse.Headers)
	assert.Equal(t, response.Body, expectedResponse.Body)
	assert.Equal(t, err, nil)
}

func TestHtml(t *testing.T) {
	request := events.APIGatewayProxyRequest{
		Headers: map[string]string{
			"accepts": "text/html",
		},
		QueryStringParameters: map[string]string{
			"to":     "Scott",
			"from":   "Kyle",
			"excuse": "dog",
		}}
	expectedResponse := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "text/html",
			"Access-Control-Allow-Origin": "*",
			"Access-Control-Allow-Method": "*",
			"Access-Control-Allow-Header": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
		},
		Body: "kyle",
	}

	response, err := Handler(request)

	assert.Equal(t, response.Headers, expectedResponse.Headers)
	assert.Contains(t, response.Body, "Scott")
	assert.Contains(t, response.Body, "Kyle")
	assert.Equal(t, err, nil)
}

func TestJson(t *testing.T) {
	request := events.APIGatewayProxyRequest{
		Headers: map[string]string{
			"accepts": "application/json",
		},
		QueryStringParameters: map[string]string{
			"to":     "Scott",
			"from":   "kyle",
			"excuse": "dog",
		}}
	expectedResponse := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Access-Control-Allow-Origin": "*",
			"Access-Control-Allow-Method": "*",
			"Access-Control-Allow-Header": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
		},
		Body: "kyle",
	}

	response, err := Handler(request)

	assert.Equal(t, response.Headers, expectedResponse.Headers)
	assert.Contains(t, response.Body, "Scott")
	assert.Contains(t, response.Body, "kyle")
	assert.Equal(t, err, nil)

}

func TestBuildInput(t *testing.T) {

}
func TestBuildMessage(t *testing.T) {

}
func TestBadInput(t *testing.T) {
	request := events.APIGatewayProxyRequest{
		Headers: map[string]string{
			"accepts": "application/json",
		},
		QueryStringParameters: map[string]string{
			"to":     "Scott",
			"from":   "kyle",
			"excuse": "notReal",
		}}
	expectedResponse := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Access-Control-Allow-Origin": "*",
			"Access-Control-Allow-Method": "*",
			"Access-Control-Allow-Header": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
		},
		Body: "kyle",
	}

	response, err := Handler(request)

	assert.Equal(t, response.Headers, expectedResponse.Headers)
	assert.Contains(t, response.Body, "Scott")
	assert.Contains(t, response.Body, "kyle")
	assert.Equal(t, err, nil)
}

func TestList(t *testing.T) {
	request := events.APIGatewayProxyRequest{
		Headers: map[string]string{
			"accepts": "application/json",
		},
		QueryStringParameters: map[string]string{
			"format": "list",

		}}
	expectedResponse := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Access-Control-Allow-Origin": "*",
			"Access-Control-Allow-Method": "*",
			"Access-Control-Allow-Header": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
		},
		Body: "kyle",
	}

	response, err := Handler(request)

	assert.Equal(t, response.Headers, expectedResponse.Headers)
	assert.Contains(t, response.Body, "message")
	assert.Contains(t, response.Body, "key")
	assert.Equal(t, err, nil)
}

