/**
* # AWS SNS Module
*
* ## Description
*
* This module is responsible for creating an Amazon SNS Topic and
* related topic policy
*
*/

resource "aws_sns_topic" "topic" {
  name = var.name
}

resource "aws_sns_topic_policy" "topic_policy" {
  policy = data.aws_iam_policy_document.topic_policy.json
  arn    = aws_sns_topic.topic.arn
}

data "aws_iam_policy_document" "topic_policy" {
  statement {
    actions = [
      "SNS:Subscribe",
      "SNS:SetTopicAttributes",
      "SNS:Receive",
      "SNS:Publish",
      "SNS:ListSubscriptionsByTopic",
      "SNS:GetTopicAttributes",
    ]

    condition {
      test     = "StringEquals"
      variable = "AWS:PrincipalOrgId"

      values = [
        var.org_id,
      ]
    }

    effect = "Allow"

    principals {
      type        = "AWS"
      identifiers = ["*"]
    }

    resources = [
      aws_sns_topic.topic.arn,
    ]

    sid = "__default_statement_ID"
  }
}
