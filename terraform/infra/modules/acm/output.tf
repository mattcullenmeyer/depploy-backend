output "acm_certificate_arn" {
  description = "The ARN of the certificate"
  value       = aws_acm_certificate.acm_certificate.arn
}

output "acm_certificate_validation_options" {
  description = "Set of domain validation objects which can be used to complete certificate validation"
  value = aws_acm_certificate.acm_certificate.domain_validation_options
}
