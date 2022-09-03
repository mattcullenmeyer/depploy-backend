terraform {
  required_version = "1.2.5"

  backend "s3" {
    bucket         = "depploy-terraform"
    key            = "api-depploy/terraform.tfstate"
    region         = "us-east-1"
    dynamodb_table = "depploy-terraform"
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}
