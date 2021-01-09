variable account_name {
  type = string
}

variable account_email {
  type = string
}

variable "region" {
  description = "Region where the resources are going to be created"
  default     = "us-east-1"
}
