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
)


//Input are the possible inputs that lambda takes
type Input struct {
	From        string
	To          string
	Excuse      string
	Format      string
	ContentType string
	request     events.APIGatewayProxyRequest
}

// Handler is executed by AWS Lambda in the main function. Once the request
// is processed, it returns an Amazon API Gateway response object to AWS Lambda
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	input := buildInput(request)

	// check if there is an error if so error out

	body, contentType, err := excuse(input)

	// check if there is an error if so error out

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       body,
		Headers: map[string]string{
			"Content-Type": contentType,
			"Access-Control-Allow-Origin": "*",
			"Access-Control-Allow-Method": "*",
			"Access-Control-Allow-Header": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
		},
	}, err
}

func buildInput(request events.APIGatewayProxyRequest) Input {

	excuse := request.QueryStringParameters["excuse"]
	to := request.QueryStringParameters["to"]
	from := request.QueryStringParameters["from"]
	contentType := request.Headers["accepts"]
	format := request.QueryStringParameters["format"]
	if request.QueryStringParameters["format"] == "slack" {

		pairs := strings.Split(request.Body, "&")

		mapped := map[string]string{}
		for _, element := range pairs {
			// element is the element from someSlice for where we are
			pairArray := strings.Split(element, "=")
			mapped[pairArray[0]] = pairArray[1]
		}

		split := strings.Split(mapped["text"], "%2C")

		if len(split) >= 1 {
			excuse = strings.TrimSpace(split[0])
		} else {
			excuse = "dog"
		}

		if excuse == "" {
			excuse = "dog"
		}
		if len(split) >= 2 {
			to = strings.TrimSpace(split[1])
		} else {
			to = ""
		}
		if len(split) >= 3 {
			from = strings.TrimSpace(split[2])
		} else {
			from = mapped["user_name"]
		}
	}
	return Input{
		From:        from,
		To:          to,
		Excuse:      excuse,
		Format:      format,
		ContentType: contentType,
		request:     request,
	}
}

func getMessage(request Input) Message {

	memo := getMemo(request)

	text := ""
	if request.To != "" {
		text += "Dear " + request.To + ", "
	}
	text += memo

	if request.From != "" {
		text += " Sincerly " + request.From
	}

	return Message{
		From: request.From,
		Memo: memo,
		To:   request.To,
		Text: text,
	}
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

func toSlack(message Message) (string, error) {
	type SlackResponse struct {
		Text string `json:"text"`
	}

	response, err := json.Marshal(SlackResponse{Text: message.Text})
	return string(response), err
}

func excuse(request Input) (string, string, error) {

	if request.Format == "list" {
		body, err := toList()
		return body, "application/json", err
	}

	message := getMessage(request)

	if request.Format == "slack" {
		body, err := toSlack(message)
		return body, "application/json", err
	}

	if request.ContentType == "application/json" {
		body, err := toJSON(message)
		return body, "application/json", err
	}
	//add plain text
	//add xml
	body, err := toHTML(message)

	return body, "text/html", err
}

func getMemo(request Input) string {

	e := mapExcuses()[request.Excuse]
	if e == nil {		
		e = listExcuses()[rand.Intn(len(listExcuses()) -1)]
	}
	return e.buildMessage()
}

func toList() (string, error) {
	type ListResponse struct {
		Doc     string `json:"doc"`
		Key     string `json:"key"`
		Message string `json:"message"`
		Params  string `json:"params"`
	}
	
	list := []ListResponse{}
	for _, e := range listExcuses() {
		list = append(list, ListResponse{
			Doc: "Doc",
			Key: e.getKey(),
			Message: e.buildMessage(),
		})
	}

	response, err := json.Marshal(list)
	return string(response), err
}

func main() {
	lambda.Start(Handler)
}
