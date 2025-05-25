/**
* # AWS ECR Module
*
* ## Description
*
* This module is responsible for creating an Amazon Elastic Container Registry (ECR)
* repository. It provides the necessary IAM policies and roles to allow access to the
* ECR repository. The module also includes a data source to retrieve the ECR repository
* URI, which can be used for pushing and pulling images.
*
*/

# Use the name variable to craft a unique name for all resources
# created by this module

data "aws_caller_identity" "current" {}

resource "aws_ecr_repository" "ecr" {
  name = var.name
}

resource "aws_ecr_repository_policy" "policy" {
  repository = aws_ecr_repository.ecr.name
  policy     = data.aws_iam_policy_document.account.json
}

data "aws_iam_policy_document" "account" {
  statement {
    sid    = "AccessFromAccount"
    effect = "Allow"

    principals {
      type        = "AWS"
      identifiers = [data.aws_caller_identity.current.account_id]
    }

    actions = [
      "ecr:GetDownloadUrlForLayer",
      "ecr:BatchGetImage",
      "ecr:BatchCheckLayerAvailability",
      "ecr:PutImage",
      "ecr:InitiateLayerUpload",
      "ecr:UploadLayerPart",
      "ecr:CompleteLayerUpload",
      "ecr:DescribeRepositories",
      "ecr:GetRepositoryPolicy",
      "ecr:ListImages",
      "ecr:DeleteRepository",
      "ecr:BatchDeleteImage",
      "ecr:SetRepositoryPolicy",
      "ecr:DeleteRepositoryPolicy",
    ]
  }
  statement {
    sid    = "LambdaPullAccess"
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }

    actions = [
      "ecr:GetDownloadUrlForLayer",
      "ecr:BatchGetImage",
      "ecr:BatchCheckLayerAvailability",
      "ecr:DescribeRepositories",
      "ecr:GetRepositoryPolicy",
      "ecr:ListImages",
    ]
  }
}
