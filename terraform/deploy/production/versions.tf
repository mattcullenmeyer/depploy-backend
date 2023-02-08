terraform {
  required_version = "~>1.3.7"

  backend "s3" {
    bucket         = "depploy-terraform-state-production"
    key            = "state/backend.tfstate"
    region         = "us-east-1"
    dynamodb_table = "depploy-terraform-lock-production"
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}
