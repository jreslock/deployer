<!-- BEGIN_TF_DOCS -->
# AWS Context Module

## Description

This module is responsible for providing various AWS contexts
use with downstream modules.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | n/a |

## Inputs

No inputs.

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_aws_account_arn"></a> [aws\_account\_arn](#output\_aws\_account\_arn) | n/a |
| <a name="output_aws_account_id"></a> [aws\_account\_id](#output\_aws\_account\_id) | n/a |
| <a name="output_aws_account_user_id"></a> [aws\_account\_user\_id](#output\_aws\_account\_user\_id) | n/a |
| <a name="output_aws_availability_zones"></a> [aws\_availability\_zones](#output\_aws\_availability\_zones) | n/a |
| <a name="output_aws_organization_accounts"></a> [aws\_organization\_accounts](#output\_aws\_organization\_accounts) | n/a |
| <a name="output_aws_organization_arn"></a> [aws\_organization\_arn](#output\_aws\_organization\_arn) | n/a |
| <a name="output_aws_organization_id"></a> [aws\_organization\_id](#output\_aws\_organization\_id) | n/a |
| <a name="output_aws_organization_master_account_arn"></a> [aws\_organization\_master\_account\_arn](#output\_aws\_organization\_master\_account\_arn) | n/a |
| <a name="output_aws_organization_master_account_name"></a> [aws\_organization\_master\_account\_name](#output\_aws\_organization\_master\_account\_name) | n/a |
| <a name="output_aws_partition"></a> [aws\_partition](#output\_aws\_partition) | n/a |
| <a name="output_aws_region"></a> [aws\_region](#output\_aws\_region) | n/a |

## Resources

| Name | Type |
|------|------|
| [aws_availability_zones.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/availability_zones) | data source |
| [aws_caller_identity.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller_identity) | data source |
| [aws_organizations_organization.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/organizations_organization) | data source |
| [aws_partition.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/partition) | data source |
| [aws_region.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/region) | data source |
<!-- END_TF_DOCS -->
