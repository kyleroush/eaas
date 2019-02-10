package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Message will not be exported but is used several places
type Message struct {
	Comment string `json:"comment"`
	From    string `json:"from"`
}

// Handler is executed by AWS Lambda in the main function. Once the request
// is processed, it returns an Amazon API Gateway response object to AWS Lambda
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Hello World")
	log.Println(request.Path)
	type Person struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}

	bytes, err := json.Marshal(Person{
		FirstName: "John",
		LastName:  "Dow",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

	// if request.Path == "" {
	// 	return mainPage(request)
	// }

	// read this from the request header

	body, contentType, err := excuse(request)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       body,
		Headers: map[string]string{
			"Content-Type": contentType,
		},
	}, err
}

func toJSON(message Message) (string, error) {
	response, err := json.Marshal(message)
	return string(response), err
}

func excuse(request events.APIGatewayProxyRequest) (string, string, error) {

	// if request.Headers["accepts"] == "text/json" {
	message, err := toJSON(getMessage())
	return message, "text/json", err
	// }
	// index, err := ioutil.ReadFile("public/index.html")

	// return string(index), "text/html", err
}

func getMessage() Message {

	return Message{From: "kyles", Comment: "was here"}
}

func mainPage(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	index, err := ioutil.ReadFile("public/index.html")
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(index),
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
	}, nil
}

func main() {
	lambda.Start(Handler)
}
