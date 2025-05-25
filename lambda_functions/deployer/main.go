package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	lambdasdk "github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

// FunctionWrapper encapsulates function actions used in the examples.
// It contains an AWS Lambda service client that is used to perform user actions.
type FunctionWrapper struct {
	LambdaClient *lambdasdk.Client
}

// Event represents the structure of the SQS message
type Event struct {
	ImageURI string `json:"image_uri"`
}

// UpdateFunctionCode updates the code for the Lambda function specified by functionName.
// The existing code for the Lambda function is entirely replaced by the code in the
// zipPackage buffer. After the update action is called, a lambda.FunctionUpdatedV2Waiter
// is used to wait until the update is successful.
// FunctionConfig stores the name and image URI of a Lambda function
type FunctionConfig struct {
	FunctionName string
	ImageURI     string
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
		return err
	}

	// Create Lambda client
	client := lambdasdk.NewFromConfig(cfg)

	for _, record := range sqsEvent.Records {
		var event Event
		if err := json.Unmarshal([]byte(record.Body), &event); err != nil {
			log.Printf("Error unmarshaling SQS message: %v", err)
			continue // Skip to the next message
		}

		log.Printf("Received image URI from SQS: %s", event.ImageURI)

		// Get all Lambda functions
		functions, err := getAllLambdaFunctions(ctx, client)
		if err != nil {
			log.Printf("Error getting all Lambda functions: %v", err)
			return err
		}

		// Iterate through Lambda functions and check image URI
		for _, function := range functions {
			if function.ImageURI != "" && event.ImageURI != "" && function.ImageURI == event.ImageURI {
				log.Printf("Found matching Lambda function: %s", function.FunctionName)

				// Update Lambda function code
				err = updateFunctionCode(ctx, client, function.FunctionName, event.ImageURI)
				if err != nil {
					log.Printf("Error updating Lambda function %s: %v", function.FunctionName, err)
					return err
				}

				log.Printf("Successfully updated Lambda function: %s", function.FunctionName)
			}
		}
	}

	return nil
}

// getAllLambdaFunctions retrieves all Lambda functions in the current AWS account and region.
func getAllLambdaFunctions(ctx context.Context, client *lambdasdk.Client) ([]FunctionConfig, error) {
	var functions []FunctionConfig
	var nextMarker *string

	for {
		output, err := client.ListFunctions(ctx, &lambdasdk.ListFunctionsInput{
			Marker: nextMarker,
		})
		if err != nil {
			return nil, fmt.Errorf("error listing functions: %w", err)
		}

		for _, function := range output.Functions {
			// Get function details to retrieve image URI
			configOutput, err := client.GetFunction(ctx, &lambdasdk.GetFunctionInput{
				FunctionName: function.FunctionName,
			})
			if err != nil {
				log.Printf("Error getting function details for %s: %v", *function.FunctionName, err)
				continue // Skip to the next function
			}

			// Extract image URI from function configuration
			imageURI := ""
			if configOutput.Configuration != nil && configOutput.Configuration.PackageType == types.PackageTypeImage && configOutput.Code != nil && configOutput.Code.ImageUri != nil {
				imageURI = *configOutput.Code.ImageUri
			}

			functions = append(functions, FunctionConfig{
				FunctionName: *function.FunctionName,
				ImageURI:     imageURI,
			})
		}

		nextMarker = output.NextMarker
		if nextMarker == nil {
			break
		}
	}

	return functions, nil
}

// updateFunctionCode updates the Lambda function code with the new image URI.
func updateFunctionCode(ctx context.Context, client *lambdasdk.Client, functionName string, imageURI string) error {
	_, err := client.UpdateFunctionCode(ctx, &lambdasdk.UpdateFunctionCodeInput{
		FunctionName: aws.String(functionName),
		ImageUri:     aws.String(imageURI),
	})
	if err != nil {
		return fmt.Errorf("error updating function code: %w", err)
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
