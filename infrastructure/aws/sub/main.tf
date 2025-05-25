/**
* # AWS Deployment Subscriber Module
*
* ## Description
*
* This module is responsible for the subscriber end of
* a deployer pipeline. It creates the Lambda, SQS and
* IAM resources required to run the deployer and receive
* events when images are pushed to ECR
*
*/

locals {
  # The publiser name is randomly generated and we don't
  # know what it will be ahead of time nor can we derive it
  # from other known inputs.
  #
  # We do have the input variable for the topic ARN which will
  # have the name we want at the end.
  name = element(split(":", var.sns_topic_arn), 5)
}

module "context" {
  source = "../../modules/aws/context"
}

module "deployer_sub" {
  source = "../../modules/aws/lambda"

  name               = "${local.name}-sub-${module.context.aws_account_id}"
  image_uri          = "${var.image_uri_prefix}/deployer_lambda:latest"
  log_retention_days = var.log_retention_days
  sns_topic_arn      = var.sns_topic_arn
}

terraform {
  backend "s3" {}
  required_version = ">=1.8.2"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">=5.97.0"
    }
    external = {
      source  = "hashicorp/external"
      version = ">=2.3.1"
    }
  }
}

provider "aws" {
  region = "us-east-1"
  default_tags {
    tags = {
      Name = local.name
    }
  }
}

provider "external" {}
