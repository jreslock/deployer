variable "image_uri_prefix" {
  description = "Image URI prefix (without the name or tag) for the Lambda Function Source Binary"
  type        = string
}

variable "log_retention_days" {
  description = "Number of days to keep lambda logs in CloudWatch"
  type        = number
}

variable "role_arn" {
  description = "AWS Role ARN for the provider to assume"
  type        = string
}

variable "sns_topic_arn" {
  description = "The SNS Topic ARN to route events to"
  type        = string
}
