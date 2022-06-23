output "name" {
  description = "Name of the virtual gateway"
  value       = module.vgw.name
}

output "domain_name" {
  description = "Domain name for the virtual gateway"
  value       = module.vgw.domain_name
}
