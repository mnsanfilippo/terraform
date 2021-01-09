output "account_id" {
  value = aws_organizations_account.dev.id
}

output "account_email" {
  value = aws_organizations_account.dev.email
}