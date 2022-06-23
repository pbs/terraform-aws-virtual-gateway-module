module "mesh" {
  source = "github.com/pbs/terraform-aws-app-mesh-module?ref=0.0.1"

  organization = var.organization
  environment  = var.environment
  product      = var.product
  repo         = var.repo
}

module "vgw" {
  source = "../.."

  primary_hosted_zone = var.primary_hosted_zone
  mesh_name           = module.mesh.name

  organization = var.organization
  environment  = var.environment
  product      = var.product
  repo         = var.repo
}
