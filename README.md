# Deployment Pipeline Infrastructure

This repository contains the infrastructure as code (IaC) for a deployment pipeline on AWS, built using OpenTofu. It includes modules for creating and managing ECR repositories, Lambda functions, SNS topics, and other related resources.

## Infrastructure Overview

The infrastructure is divided into two main parts:

* **Publisher:** Responsible for publishing ECR push events to an SNS topic.
* **Subscriber:** Responsible for subscribing to the SNS topic and triggering a Lambda function when new images are pushed to ECR.

### Modules

The following OpenTofu modules are used to create the infrastructure:

* **[modules/aws/context](infrastructure/modules/aws/context/README.md):** Provides AWS context information such as account ID, region, and partition.
* **[modules/aws/ecr](infrastructure/modules/aws/ecr/README.md):** Creates and manages an ECR repository for storing container images.
* **[modules/aws/lambda](infrastructure/modules/aws/lambda/README.md):** Creates and manages a Lambda function for processing events.
* **[modules/aws/sns](infrastructure/modules/aws/sns/README.md):** Creates and manages an SNS topic for publishing events.

### Components

The following components are created as part of the infrastructure:

* ECR Repository: Stores the container images for the Lambda function.
* Lambda Function: Processes events triggered by ECR push events.
* SNS Topic: Publishes ECR push events to subscribers.
* CloudWatch Event Bus: Routes events from ECR to the SNS topic.
* CloudWatch Event Rule: Filters events from ECR based on the event type.
* IAM Roles and Policies: Provide the necessary permissions for the Lambda function and other resources to access AWS services.
* SQS Queues: Provides queueing for the lambda function.

### Publisher ([infrastructure/aws/pub](infrastructure/aws/pub/README.md))

The publisher component is responsible for publishing ECR push events to an SNS topic. It includes the following resources:

* ECR repository (`deployer_lambda`) to store function's container image.
* SNS topic to send notifications about image pushes.
* CloudWatch event bus, rule and target to route ECR push events to the SNS topic.

### Subscriber ([infrastructure/aws/sub](infrastructure/aws/sub/README.md))

The subscriber component is responsible for subscribing to the SNS topic and triggering a Lambda function when new images are pushed to ECR. It includes the following resources:

* Lambda function to process events.
* SQS queue to queue events for processing.
* IAM roles and policies to provide the necessary permissions for the Lambda function to access AWS services.

## Getting Started

To get started with this infrastructure, you will need the following:

* An AWS account
* OpenTofu installed
* The AWS CLI installed and configured

## Deployment

To deploy the infrastructure, follow these steps:

1. Clone the repository.
2. Navigate to the `infrastructure/aws/pub` directory.
3. Create `target.tfvars` and `backend.vars` files with the required variables.
4. Initialize OpenTofu: `tofu init`
5. Apply the OpenTofu configuration: `tofu apply`
6. Navigate to the `infrastructure/aws/sub` directory.
7. Create `target.tfvars` and `backend.vars` files with the required variables, including the `sns_topic_arn` from the publisher deployment.
8. Initialize OpenTofu: `tofu init`
9. Apply the OpenTofu configuration: `tofu apply`

## Contributing

Contributions are welcome! Please fork this repository and submit a pull request with your changes against the `main` branch.

## License

This project is licensed under the MIT License.
