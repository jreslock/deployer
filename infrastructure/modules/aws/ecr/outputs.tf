output "ecr_arn" {
  description = "ARN of the ECR repository"
  value       = aws_ecr_repository.ecr.arn
}
