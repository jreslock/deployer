variable "name" {
  description = "Name for all resources created by this module"
  type        = string
}

variable "image_uri" {
  description = "ECR Image URI for the container image containing the lambda function code"
  type        = string
}

variable "log_retention_days" {
  description = "Number of days to keep logs for the lambda function"
  type        = number
}

variable "sns_topic_arn" {
  description = "SNS Topic ARN to allow sending messages to the SQS Queue"
  type        = string
}
