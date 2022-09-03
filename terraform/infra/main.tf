module "aws_acm" {
  source      = "./modules/acm"
  domain_name = var.domain_name
}
