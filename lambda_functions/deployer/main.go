package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	lambdasdk "github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

// Event represents the structure of the SQS message
type Event struct {
	ImageURI string `json:"image_uri"`
}

// FunctionConfig stores the name and image URI of a Lambda function
type FunctionConfig struct {
	FunctionName string
	ImageURI     string
}

type LambdaAPI interface {
	ListFunctions(ctx context.Context, params *lambdasdk.ListFunctionsInput, optCns ...func(*lambdasdk.Options)) (*lambdasdk.ListFunctionsOutput, error)
	GetFunction(ctx context.Context, params *lambdasdk.GetFunctionInput, optFns ...func(*lambdasdk.Options)) (*lambdasdk.GetFunctionOutput, error)
}

type LambdaUpdateFunctionCodeAPI interface {
	UpdateFunctionCode(ctx context.Context, params *lambdasdk.UpdateFunctionCodeInput, optFns ...func(*lambdasdk.Options)) (*lambdasdk.UpdateFunctionCodeOutput, error)
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		slog.Error("failed to load configuration", slog.Any("error", err))
		return err
	}

	// Create Lambda client
	client := lambdasdk.NewFromConfig(cfg)

	for _, record := range sqsEvent.Records {
		var event Event
		if err := json.Unmarshal([]byte(record.Body), &event); err != nil {
			slog.Warn("Error unmarshaling SQS message", slog.Any("error", err))
			continue // Skip to the next message
		}

		slog.Info("Received image URI from SQS", slog.String("imageURI", event.ImageURI))

		// Get all Lambda functions
		functions, err := getAllLambdaFunctions(ctx, client)
		if err != nil {
			slog.Error("Error getting all Lambda functions", slog.Any("error", err))
			return err
		}

		// Iterate through Lambda functions and check image URI
		for _, function := range functions {
			if function.ImageURI != "" && event.ImageURI != "" && function.ImageURI == event.ImageURI {
				slog.Info("Found matching Lambda function", slog.String("functionName", function.FunctionName))

				// Update Lambda function code
				err = updateFunctionCode(ctx, client, function.FunctionName, event.ImageURI)
				if err != nil {
					slog.Error("Error updating Lambda function", slog.String("functionName", function.FunctionName), slog.Any("error", err))
					return err
				}

				slog.Info("Successfully updated Lambda function", slog.String("functionName", function.FunctionName))
			}
		}
	}

	return nil
}

// getAllLambdaFunctions retrieves all Lambda functions in the current AWS account and region.
func getAllLambdaFunctions(ctx context.Context, client LambdaAPI) ([]FunctionConfig, error) {
	functions := []FunctionConfig{}
	var nextMarker *string

	for {
		output, err := client.ListFunctions(ctx, &lambdasdk.ListFunctionsInput{
			Marker: nextMarker,
		})
		if err != nil {
			return nil, err
		}

		for _, function := range output.Functions {
			// Get function details to retrieve image URI
			configOutput, err := client.GetFunction(ctx, &lambdasdk.GetFunctionInput{
				FunctionName: function.FunctionName,
			})
			if err != nil {
				slog.Warn("Error getting function details", slog.String("functionName", *function.FunctionName), slog.Any("error", err))
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
func updateFunctionCode(ctx context.Context, client LambdaUpdateFunctionCodeAPI, functionName, imageURI string) error {
	input := &lambdasdk.UpdateFunctionCodeInput{
		FunctionName: aws.String(functionName),
		ImageUri:     aws.String(imageURI),
		Publish:      true,
	}
	_, err := client.UpdateFunctionCode(ctx, input)

	return err
}

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})))
	lambda.Start(handler)
}
