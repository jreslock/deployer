output "aws_region" {
  value = data.aws_region.current.name
}
output "aws_partition" {
  value = data.aws_partition.current.partition
}
output "aws_availability_zones" {
  value = data.aws_availability_zones.current.names
}
output "aws_account_id" {
  value = data.aws_caller_identity.current.account_id
}

output "aws_account_arn" {
  value = data.aws_caller_identity.current.arn
}
output "aws_account_user_id" {
  value = data.aws_caller_identity.current.user_id
}

output "aws_organization_id" {
  value = data.aws_organizations_organization.current.id
}
output "aws_organization_arn" {
  value = data.aws_organizations_organization.current.arn
}

output "aws_organization_master_account_arn" {
  value = data.aws_organizations_organization.current.master_account_arn
}
output "aws_organization_master_account_name" {
  value = data.aws_organizations_organization.current.master_account_name
}
output "aws_organization_accounts" {
  value = data.aws_organizations_organization.current.accounts
}
