data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

resource "aws_iam_role" "lambda" {
  name               = var.name
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

resource "aws_iam_role_policy" "lambda_logging" {
  role   = aws_iam_role.lambda.name
  policy = data.aws_iam_policy_document.lambda_logging.json
}

resource "aws_iam_role_policy" "lambda_work_requirements" {
  role   = aws_iam_role.lambda.name
  policy = data.aws_iam_policy_document.lambda_work_requirements.json
}

data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}

data "aws_iam_policy_document" "lambda_logging" {
  statement {
    effect = "Allow"

    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]

    resources = ["${aws_cloudwatch_log_group.log_group.arn}:*"]
  }
}

data "aws_iam_policy_document" "lambda_work_requirements" {
  statement {
    effect = "Allow"

    actions = [
      "lambda:UpdateFunctionCode",
      "lambda:GetAlias",
      "lambda:GetFunction",
      "lambda:ListAliases",
      "lambda:ListFunctions",
      "lambda:ListTags",
      "lambda:PublishVersion",
      "lambda:TagResource",
      "lambda:UpdateAlias"
    ]

    resources = ["arn:aws:lambda:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:function:*"]
  }
  statement {
    effect = "Allow"

    actions = [
      "sqs:ReceiveMessage",
      "sqs:DeleteMessage",
      "sqs:GetQueueAttributes"
    ]

    resources = [aws_sqs_queue.queue.arn]
  }
}
