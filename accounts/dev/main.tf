resource "aws_organizations_account" "mnsanfilippo-dev" {
  name  = var.account_name
  email = var.account_email
}

