package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"strings"
	"text/template"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kyleroush/eaas/excuses"
	"github.com/kyleroush/eaas/models"
)

// Message This is the response of the service
type Message struct {
	Text string `json:"text"`
}

// Handler is executed by AWS Lambda in the main function. Once the request
// is processed, it returns an Amazon API Gateway response object to AWS Lambda
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	input := buildInput(request)

	// check if there is an error if so error out

	body, contentType, err := excuse(input)

	// check if there is an error if so error out

	// todo check if teh access headers are still required
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       body,
		Headers: map[string]string{
			"Content-Type":                contentType,
			"Access-Control-Allow-Origin": "*",
			"Access-Control-Allow-Method": "GET",
			"Access-Control-Allow-Header": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
		},
	}, err
}

// builds the input for the service
func buildInput(request events.APIGatewayProxyRequest) models.Input {

	inputs := map[string]string{}

	for k, v := range request.QueryStringParameters {
		inputs[k] = v
	}

	contentType := request.Headers["accepts"]
	format := request.QueryStringParameters["format"]

	// basicly if the format is slack that means the params will be coming from a multi part form
	// so we need to read the params differently by parsing them out of the text field
	if request.QueryStringParameters["format"] == "slack" {

		pairs := strings.Split(request.Body, "&")

		// map the values so that we can look up text
		mapped := map[string]string{}
		for _, element := range pairs {
			// element is the element from someSlice for where we are
			pairArray := strings.Split(element, "=")
			mapped[pairArray[0]] = pairArray[1]
		}

		// todo: change this to not be comma seperate be equals
		split := strings.Split(mapped["text"], "%2C")

		for _, element := range split {
			// element is the element from someSlice for where we are
			pairArray := strings.Split(element, "=")
			inputs[pairArray[0]] = pairArray[1]
		}

		// todo decide if i still want to use the user_name as an optional value
		// } else {
		// 	from = mapped["user_name"]
		// }
	}
	return models.Input{
		Format:      format,
		ContentType: contentType,
		Request:     request,
		Inputs:      inputs,
	}
}

// toHTML: create teh response in the form of html
func toHTML(message Message) (string, error) {
	t, err := template.ParseFiles("public/message.html") //parse the html file homepage.html
	if err != nil {                                      // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	w := bytes.NewBufferString("")
	err = t.Execute(w, message) //execute the template and pass it the HomePageVars struct to fill in the gaps
	return w.String(), nil
}

// toJSON build the response in the format of a json
func toJSON(message Message) (string, error) {
	response, err := json.Marshal(message)
	return string(response), err
}

// todo: rename this method
// excuse: get the creates the response values that are dynamic
func excuse(request models.Input) (string, string, error) {

	// get list from a differet value make a specfic one for it
	if request.Format == "list" {
		body, err := toList()
		return body, "application/json", err
	}

	message := Message{
		Text: getText(request),
	}

	if request.ContentType == "application/json" {
		body, err := toJSON(message)
		return body, "application/json", err
	}
	//add plain text
	//add xml

	// by default respond with html
	body, err := toHTML(message)
	return body, "text/html", err
}

// getText: give an input look up and return the message of one of the excuses
// if there is no key Excuse key choose a random one
func getText(request models.Input) string {

	e := excuses.MapExcuses()[request.Inputs["key"]]
	if e == nil {
		e = excuses.ListExcuses()[rand.Intn(len(excuses.ListExcuses())-1)]
	}
	return e.BuildText(request)
}

// toList: builds a string witha json array of all excuses
func toList() (string, error) {
	type ListResponse struct {
		Doc    string `json:"doc"`
		Key    string `json:"key"`
		Text   string `json:"text"`
		Params string `json:"params"`
	}

	list := []ListResponse{}
	for _, e := range excuses.ListExcuses() {
		list = append(list, ListResponse{
			Doc: e.GetDoc(),
			Key: e.GetKey(),
		})
	}

	response, err := json.Marshal(list)
	return string(response), err
}

func main() {
	lambda.Start(Handler)
}
