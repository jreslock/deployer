<!-- BEGIN_TF_DOCS -->
# AWS Deployment Subscriber Module

## Description

This module is responsible for the subscriber end of
a deployer pipeline. It creates the Lambda, SQS and
IAM resources required to run the deployer and receive
events when images are pushed to ECR

## Providers

No providers.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_image_uri_prefix"></a> [image\_uri\_prefix](#input\_image\_uri\_prefix) | Image URI prefix (without the name or tag) for the Lambda Function Source Binary | `string` | n/a | yes |
| <a name="input_log_retention_days"></a> [log\_retention\_days](#input\_log\_retention\_days) | Number of days to keep lambda logs in CloudWatch | `number` | n/a | yes |
| <a name="input_sns_topic_arn"></a> [sns\_topic\_arn](#input\_sns\_topic\_arn) | The SNS Topic ARN to route events to | `string` | n/a | yes |

## Outputs

No outputs.

## Resources

No resources.
<!-- END_TF_DOCS -->
