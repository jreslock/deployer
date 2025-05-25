variable "name" {
  description = "Name for all resources created by this module"
  type        = string
}

variable "org_id" {
  description = "AWS Organization ID to allow in the SNS topic policy"
  type        = string
}
