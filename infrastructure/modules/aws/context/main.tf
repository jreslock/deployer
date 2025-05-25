/**
* # AWS Context Module
*
* ## Description
*
* This module is responsible for providing various AWS contexts
* use with downstream modules.
*
*/

data "aws_caller_identity" "current" {}
data "aws_region" "current" {}
data "aws_partition" "current" {}
data "aws_availability_zones" "current" {}
data "aws_organizations_organization" "current" {}
