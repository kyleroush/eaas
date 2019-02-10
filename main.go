package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"text/template"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func listExcuses() map[string]string {

	return map[string]string{
		"dog": "My dog ate my homework",
	}
}

// Message will not be exported but is used several places
type Message struct {
	Memo string `json:"memo"`
	From string `json:"from"`
	To   string `json:"to"`
}

// Handler is executed by AWS Lambda in the main function. Once the request
// is processed, it returns an Amazon API Gateway response object to AWS Lambda
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if request.QueryStringParameters["excuse"] == "" {
		return mainPage(request)
	}
	//if key list return all

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

func toHTML(message Message) (string, error) {
	t, err := template.ParseFiles("public/message.html") //parse the html file homepage.html
	if err != nil {                                      // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	w := bytes.NewBufferString("")
	err = t.Execute(w, message) //execute the template and pass it the HomePageVars struct to fill in the gaps
	return w.String(), nil
}

func toJSON(message Message) (string, error) {
	response, err := json.Marshal(message)
	return string(response), err
}

func excuse(request events.APIGatewayProxyRequest) (string, string, error) {

	message := getMessage(request)

	if request.Headers["Accept"] == "application/json" {
		message, err := toJSON(message)
		return message, "application/json", err
	}
	//add plain text
	//add xml
	body, err := toHTML(message)

	return body, "text/html", err
}

func getMessage(request events.APIGatewayProxyRequest) Message {

	memo := getMemo(request)

	return Message{
		From: request.QueryStringParameters["from"],
		Memo: memo,
		To:   request.QueryStringParameters["to"]}
}

func getMemo(request events.APIGatewayProxyRequest) string {
	return listExcuses()[request.QueryStringParameters["excuse"]]
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
