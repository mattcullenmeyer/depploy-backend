variable "root_domain" {
  type        = string
  description = "Exact domain for SSL/TLS (eg depploy.io)"
}

variable "subdomain" {
  type        = string
  default     = ""
  description = "Subdomain, excluding root domain (eg app)"
}
