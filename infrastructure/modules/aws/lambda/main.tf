/**
* # AWS Lambda Module
*
* ## Description
*
* This module is responsible for creating a Amazon Lambda function
* and related resources for triggering from SQS.
*
*/

# Function
resource "aws_lambda_function" "function" {
  function_name = var.name
  image_uri     = var.image_uri
  package_type  = "Image"
  role          = aws_iam_role.lambda.arn

  lifecycle {
    ignore_changes = [image_uri]
  }

  depends_on = [aws_cloudwatch_log_group.log_group]
}

# Alias
resource "aws_lambda_alias" "alias" {
  name             = var.name
  description      = "Alias for ${var.name} function"
  function_name    = aws_lambda_function.function.function_name
  function_version = "$LATEST"
}

resource "aws_lambda_permission" "invoke" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.function.function_name
  principal     = "sqs.amazonaws.com"
  source_arn    = aws_sqs_queue.queue.arn
  qualifier     = aws_lambda_alias.alias.name
}

# Log group
resource "aws_cloudwatch_log_group" "log_group" {
  name              = "/aws/lambda/${var.name}"
  retention_in_days = var.log_retention_days
}

# Event source from SQS
resource "aws_lambda_event_source_mapping" "event_source_mapping" {
  event_source_arn = aws_sqs_queue.queue.arn
  enabled          = true
  function_name    = aws_lambda_function.function.arn
  batch_size       = 1
}
