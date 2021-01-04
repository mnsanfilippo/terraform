resource "aws_organizations_account" "dev" {
  name  = var.account_name
  email = var.account_email
}

