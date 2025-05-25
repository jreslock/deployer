variable "image_publish_role_arn" {
  description = "Role ARN of the publisher"
  type        = string
}

variable "name" {
  description = "Name for the ECR Repository to be created"
  type        = string
}
