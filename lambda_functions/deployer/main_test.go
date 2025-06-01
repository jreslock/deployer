package main

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	lambdasdk "github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

// MockLambdaClient is a mock implementation of the Lambda client for testing.
type MockLambdaClient struct {
	ListFunctionsFunc               func(ctx context.Context, params *lambdasdk.ListFunctionsInput, optFns ...func(*lambdasdk.Options)) (*lambdasdk.ListFunctionsOutput, error)
	GetFunctionFunc                 func(ctx context.Context, params *lambdasdk.GetFunctionInput, optFns ...func(*lambdasdk.Options)) (*lambdasdk.GetFunctionOutput, error)
	UpdateFunctionCodeFunc          func(ctx context.Context, params *lambdasdk.UpdateFunctionCodeInput, optFns ...func(*lambdasdk.Options)) (*lambdasdk.UpdateFunctionCodeOutput, error)
	UpdateFunctionConfigurationFunc func(ctx context.Context, params *lambdasdk.UpdateFunctionConfigurationInput, optFns ...func(*lambdasdk.Options)) (*lambdasdk.UpdateFunctionConfigurationOutput, error)
}

func (m *MockLambdaClient) ListFunctions(ctx context.Context, params *lambdasdk.ListFunctionsInput, optFns ...func(*lambdasdk.Options)) (*lambdasdk.ListFunctionsOutput, error) {
	if m.ListFunctionsFunc != nil {
		return m.ListFunctionsFunc(ctx, params, optFns...)
	}
	return &lambdasdk.ListFunctionsOutput{}, nil
}

func (m *MockLambdaClient) GetFunction(ctx context.Context, params *lambdasdk.GetFunctionInput, optFns ...func(*lambdasdk.Options)) (*lambdasdk.GetFunctionOutput, error) {
	if m.GetFunctionFunc != nil {
		return m.GetFunctionFunc(ctx, params, optFns...)
	}
	return &lambdasdk.GetFunctionOutput{}, nil
}

func (m *MockLambdaClient) UpdateFunctionCode(ctx context.Context, params *lambdasdk.UpdateFunctionCodeInput, optFns ...func(*lambdasdk.Options)) (*lambdasdk.UpdateFunctionCodeOutput, error) {
	if m.UpdateFunctionCodeFunc != nil {
		return m.UpdateFunctionCodeFunc(ctx, params, optFns...)
	}
	return &lambdasdk.UpdateFunctionCodeOutput{}, nil
}

func (m *MockLambdaClient) UpdateFunctionConfiguration(ctx context.Context, params *lambdasdk.UpdateFunctionConfigurationInput, optFns ...func(*lambdasdk.Options)) (*lambdasdk.UpdateFunctionConfigurationOutput, error) {
	if m.UpdateFunctionConfigurationFunc != nil {
		return m.UpdateFunctionConfigurationFunc(ctx, params, optFns...)
	}
	return &lambdasdk.UpdateFunctionConfigurationOutput{}, nil
}

func TestHandler(t *testing.T) {
	testCases := []struct {
		name          string
		sqsEvent      events.SQSEvent
		functions     []types.FunctionConfiguration
		expectedError error
	}{
		{
			name: "Successful update",
			sqsEvent: events.SQSEvent{
				Records: []events.SQSMessage{
					{
						Body: `{"image_uri": "test-image:latest"}`,
					},
				},
			},
			functions: []types.FunctionConfiguration{
				{
					FunctionName: aws.String("test-function"),
				},
			},
			expectedError: nil,
		},
		{
			name: "Error unmarshaling SQS message",
			sqsEvent: events.SQSEvent{
				Records: []events.SQSMessage{
					{
						Body: "invalid json",
					},
				},
			},
			expectedError: nil, // Expect no error to be returned, just a log message.
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
		})
	}
}

func TestGetAllLambdaFunctions(t *testing.T) {
	testCases := []struct {
		name          string
		functions     []types.FunctionConfiguration
		expected      []FunctionConfig
		expectedError error
	}{
		{
			name: "Successful retrieval",
			functions: []types.FunctionConfiguration{
				{
					FunctionName: aws.String("test-function"),
				},
			},
			expected: []FunctionConfig{
				{
					FunctionName: "test-function",
					ImageURI:     "test-image:latest",
				},
			},
			expectedError: nil,
		},
		{
			name:          "No functions",
			functions:     []types.FunctionConfiguration{},
			expected:      []FunctionConfig{},
			expectedError: nil,
		},
		{
			name:          "Error listing functions",
			expectedError: errors.New("error listing functions"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.TODO()

			mockClient := &MockLambdaClient{
				ListFunctionsFunc: func(ctx context.Context, params *lambdasdk.ListFunctionsInput, optFns ...func(*lambdasdk.Options)) (*lambdasdk.ListFunctionsOutput, error) {
					if tc.expectedError != nil && tc.expectedError.Error() == "error listing functions" {
						return nil, errors.New("error listing functions")
					}
					functionSummaries := append([]types.FunctionConfiguration{}, tc.functions...)

					return &lambdasdk.ListFunctionsOutput{
						Functions: functionSummaries,
					}, nil
				},
				GetFunctionFunc: func(ctx context.Context, params *lambdasdk.GetFunctionInput, optFns ...func(*lambdasdk.Options)) (*lambdasdk.GetFunctionOutput, error) {
					// Only return function details if there are functions to retrieve
					if len(tc.functions) > 0 {
						functionOutput := &lambdasdk.GetFunctionOutput{
							Configuration: &types.FunctionConfiguration{
								FunctionName: aws.String("test-function"),
								PackageType:  types.PackageTypeImage,
							},
							Code: &types.FunctionCodeLocation{
								ImageUri: aws.String("test-image:latest"),
							},
						}
						return functionOutput, nil
					}
					return &lambdasdk.GetFunctionOutput{}, nil
				},
			}

			functions, err := getAllLambdaFunctions(ctx, mockClient)

			if tc.expectedError != nil {
				if err == nil {
					t.Fatalf("Expected error: %v, but got nil", tc.expectedError.Error())
				}
				if err.Error() != tc.expectedError.Error() {
					t.Fatalf("Expected error: %v, but got: %v", tc.expectedError, err)
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}

				if !reflect.DeepEqual(functions, tc.expected) {
					if len(tc.expected) > 0 && len(functions) > 0 && len(tc.expected) == len(functions) && tc.expected[0].ImageURI != functions[0].ImageURI {
						t.Fatalf("Expected: %v, but got: %v", tc.expected[0].ImageURI, functions[0].ImageURI)
					} else {
						t.Fatalf("Expected: %v, but got: %v", tc.expected, functions)
					}
				}
			}
		})
	}
}

func TestUpdateFunctionCode(t *testing.T) {
	testCases := []struct {
		name          string
		functionName  string
		imageURI      string
		expectedError error
	}{
		{
			name:          "Successful update",
			functionName:  "test-function",
			imageURI:      "test-image:latest",
			expectedError: nil,
		},
		{
			name:          "Update fails",
			functionName:  "test-function",
			imageURI:      "test-image:latest",
			expectedError: errors.New("update failed"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.TODO()

			mockClient := &MockLambdaClient{
				UpdateFunctionCodeFunc: func(ctx context.Context, params *lambdasdk.UpdateFunctionCodeInput, optFns ...func(*lambdasdk.Options)) (*lambdasdk.UpdateFunctionCodeOutput, error) {
					if tc.expectedError != nil && tc.expectedError.Error() == "update failed" {
						return nil, errors.New("update failed")
					}
					return &lambdasdk.UpdateFunctionCodeOutput{}, nil
				},
			}

			err := updateFunctionCode(ctx, mockClient, tc.functionName, tc.imageURI)

			if tc.expectedError != nil {
				if err == nil {
					t.Fatalf("Expected error: %v, but got nil", tc.expectedError)
				}
				if err.Error() != tc.expectedError.Error() {
					t.Fatalf("Expected error: %v, but got: %v", tc.expectedError, err)
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
			}
		})
	}
}
