module "account" {
  source = "git@github.com:mnsanfilippo/terraform-modules/organization/accounts.git"
  account_name = var.account_name
  account_email = var.account_email
}