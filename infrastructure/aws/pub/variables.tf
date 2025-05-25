variable "name" {
  description = "Name for all resources created"
  type        = string
}

variable "role_arn" {
  description = "AWS Role ARN for the provider to assume"
  type        = string
}
