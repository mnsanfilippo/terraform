module "account" {
  source = "git@github.com:mnsanfilippo/terraform-modules.git//aws/organization/accounts"
  account_name = var.account_name
  account_email = var.account_email
}
