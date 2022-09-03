# https://registry.terraform.io/providers/hashicorp%20%20/aws/latest/docs/resources/acm_certificate
resource "aws_acm_certificate" "acm_certificate" {
  domain_name       = var.domain_name
  validation_method = "DNS"

  lifecycle {
    create_before_destroy = true
  }
}

# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/acm_certificate_validation
resource "aws_acm_certificate_validation" "aws_acm_certificate_validation" {
  certificate_arn = aws_acm_certificate.acm_certificate.arn
}
