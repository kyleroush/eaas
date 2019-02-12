package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
	"text/template"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func listExcuses() map[string]string {

	return map[string]string{
		"dog": "My dog ate my homework",
	}
}

//Input are the possible inputs that lambda takes
type Input struct {
	From        string
	To          string
	Excuse      string
	Format      string
	ContentType string
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

	log.Print(request.Body)
	// if request.QueryStringParameters["excuse"] == "" {
	// 	return mainPage(request)
	// }
	//if key list return all

	// read this from the request header

	input := buildInput(request)

	body, contentType, err := excuse(input)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       body,
		Headers: map[string]string{
			"Content-Type": contentType,
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

		type Slack struct {
			Text     string `json:"text"`
			UserName string `json:"user_name"`
		}

		data := &Slack{}
		json.Unmarshal([]byte(request.Body), data)

		split := strings.Split(data.Text, ",")

		excuse = strings.TrimSpace(split[0])
		to = strings.TrimSpace(split[1])
		from = strings.TrimSpace(split[2])

		if from == "" {
			from = data.UserName
		}
	}
	return Input{
		From:        from,
		To:          to,
		Excuse:      excuse,
		Format:      format,
		ContentType: contentType,
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
	response, err := json.Marshal(SlackResponse{Text: "Dear " + message.To + ", " + message.Memo + " Sincerly" + message.From})
	return string(response), err
}

func excuse(request Input) (string, string, error) {

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

func getMessage(request Input) Message {

	memo := getMemo(request)
	return Message{
		From: request.From,
		Memo: memo,
		To:   request.To}
}

func getMemo(request Input) string {
	return listExcuses()[request.Excuse]
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
