variable "acm_certificate_validation_options" {
  type        = any
  description = "Set of domain validation objects which can be used to complete certificate validation"
}

variable "hosted_zone_id" {
  type        = string
  description = "Route 53 hosted zone ID"
}
