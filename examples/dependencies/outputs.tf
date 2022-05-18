output "service_network_spoke" {
  value = {
    "network_name"        = values(module.service_network_spoke)[*].network_name
    "resource_group_name" = values(module.resource_group_spoke)[*].name
  }
}

output "service_network_hub" {
  value = {
    "network_name"        = values(module.service_network_hub)[*].network_name
    "resource_group_name" = values(module.resource_group_hub)[*].name
  }
}