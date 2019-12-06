// +build !test

package main

import "github.com/aws/aws-lambda-go/lambda"

func main() {
	// boostrap proper transferer implementation from your domain package here.
	var transferer Transferer
	lambda.Start(NewHandler(transferer))
}
