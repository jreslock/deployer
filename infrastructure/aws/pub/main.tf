/**
* # AWS Deployment Publisher Module
*
* ## Description
*
* This module is responsible for the publisher end of
* a deployer pipeline. It creates the ECR, Event Bus,
* Rule, Target and SNS Topic for publishing ECR push
* events to SNS. The SNS Topic has an associated policy
* allowing widely-scoped subscription permission so many
* subscribers can source from a single publisher if desired.
*
*/

# Context about our AWS connection
module "context" {
  source = "../../modules/aws/context"
}

# A random ID to generate uniqueness in names
resource "random_id" "suffix" {
  byte_length = 4
}

# Use the random ID to make a unique name
locals {
  name = "${var.name}-${random_id.suffix.dec}"
}

# An ECR for the function's container image
module "ecr" {
  source                 = "../../modules/aws/ecr"
  name                   = "deployer_lambda"
  image_publish_role_arn = "arn:aws:iam::${module.context.aws_account_id}:role/github_actions"
}

# Create an inline IAM policy for the actions role so it can
# push to ECR
resource "aws_iam_role_policy" "actions_ecr_publish" {
  role   = "github_actions"
  policy = data.aws_iam_policy_document.actions_ecr_publish.json
}

data "aws_iam_policy_document" "actions_ecr_publish" {
  statement {
    sid    = "AllowECRPublishing"
    effect = "Allow"

    resources = [module.ecr.ecr_arn]

    actions = [
      "ecr:CompleteLayerUpload",
      "ecr:UploadLayerPart",
      "ecr:InitiateLayerUpload",
      "ecr:BatchCheckLayerAvailability",
      "ecr:PutImage",
      "ecr:BatchGetImage"
    ]
  }
  statement {
    sid    = "GetECRAuthToken"
    effect = "Allow"

    resources = ["*"]

    actions = [
      "ecr:GetAuthorizationToken"
    ]
  }
}

# An SNS topic to send event notifications to
module "sns" {
  source = "../../modules/aws/sns"
  name   = local.name
  org_id = module.context.aws_organization_id
}

# An Event bus to route and filter events
resource "aws_cloudwatch_event_bus" "deployer_bus" {
  name = local.name
}

# A rule to trigger on events we care about
resource "aws_cloudwatch_event_rule" "console" {
  name           = "capture-ecr-push-events"
  description    = "Capture all ECR push events"
  event_bus_name = aws_cloudwatch_event_bus.deployer_bus.name
  event_pattern = jsonencode({
    source = [
      "aws.ecr"
    ]
    detail = {
      action-type = ["PUSH"]
      result      = ["SUCCESS"]
    }
  })
}

# A target to send events to our topic
resource "aws_cloudwatch_event_target" "sns" {
  rule           = aws_cloudwatch_event_rule.console.name
  event_bus_name = aws_cloudwatch_event_bus.deployer_bus.name
  target_id      = "SendToSNS"
  arn            = module.sns.topic_arn
}

# Provider/boilerplate to set up the providers
provider "aws" {
  region = "us-east-1"
  default_tags {
    tags = {
      Name = local.name
    }
  }
}

provider "external" {}

provider "random" {}

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
    random = {
      source = "hashicorp/random"
    }
  }
}
