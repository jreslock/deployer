<!-- BEGIN_TF_DOCS -->
# AWS Deployment Publisher Module

## Description

This module is responsible for the publisher end of
a deployer pipeline. It creates the ECR, Event Bus,
Rule, Target and SNS Topic for publishing ECR push
events to SNS. The SNS Topic has an associated policy
allowing widely-scoped subscription permission so many
subscribers can source from a single publisher if desired.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 5.98.0 |
| <a name="provider_random"></a> [random](#provider\_random) | 3.7.2 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_name"></a> [name](#input\_name) | Name for all resources created | `string` | n/a | yes |
| <a name="input_role_arn"></a> [role\_arn](#input\_role\_arn) | AWS Role ARN for the provider to assume | `string` | n/a | yes |

## Outputs

No outputs.

## Resources

| Name | Type |
|------|------|
| [aws_cloudwatch_event_bus.deployer_bus](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_event_bus) | resource |
| [aws_cloudwatch_event_rule.console](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_event_rule) | resource |
| [aws_cloudwatch_event_target.sns](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_event_target) | resource |
| [random_id.suffix](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/id) | resource |
<!-- END_TF_DOCS -->
