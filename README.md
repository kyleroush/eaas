Excuse as a Service
===================

When there is no excuse for not having an excuse

Please see [todo] for API documentation and examples.

# Development

## Install golang

## Install dependencies
*Todo* set up dependecy managment

  - `go get golang.org/x/lint/golint`
  - `go get github.com/stretchr/testify`
  - `go get github.com/aws/aws-lambda-go/events`
  - `go get github.com/aws/aws-lambda-go/lambda`
## Test
  - go lint
  - go test

# Contributing

## Adding new operations

To add a new excuse:

1. Fork into your account.
2. Create an excuse in `/excuses.
    * Use a pre built struct like Simple excuse. 
    interface: https://github.com/kyleroush/eaas/blob/master/models/simple_excuse.go
    example: https://github.com/kyleroush/eaas/blob/master/excuses/contagious_excuse.go#L5
    * Build your own custom strcut extending the excuse intereface. https://github.com/kyleroush/eaas/blob/master/models/excuse_interface.go
3. Add your new excuse to the list of excuses. https://github.com/kyleroush/eaas/blob/master/excuses/excuses.go#L7
4. Add test.
    1. Test your new excuse.
    2. Test that MapExcuses has your excuses key mapped to your excuses (we do this make sure no other excuse has a conflicting key).

All contributions are very welcome.

# Usages

## Clients
none yet but always welcome


| Language   | Reference | Info |
|:-----------|:----------|:-----|
| javascript |           | Todo |

## Integrations

| Tool  | Reference | Info |
|:------|:----------|:-----|
| Slack |In progress| TODO |
