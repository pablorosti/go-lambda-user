package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pablorosti/go-lambda-user/awsgo"
)

func main() {
	lambda.Start(EjecuteLambda)
}

func EjecuteLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InitializeAws()

	if !ValidateParams() {
		fmt.Println("Error params")

		err := errors.New("Error params")

		return event, err
	}

}

func ValidateParams() bool {
	var withParams bool
	_, withParams = os.LookupEnv("SecretName")
	return withParams
}
