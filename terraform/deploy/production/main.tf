locals {
  root_domain = "depploy.io"
  subdomain   = "api"
}

provider "aws" {
  region = "us-east-1"
}

module "infrastructure" {
  source      = "../../infra"
  root_domain = local.root_domain
  subdomain   = local.subdomain
}
