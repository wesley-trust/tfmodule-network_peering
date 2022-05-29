module "network_peering_spoke" {
  for_each                   = var.service_location
  source                     = "../"
  service_environment        = var.service_environment
  resource_network_peer      = var.service_network_hub[each.value].network_name
  resource_group_peer        = var.service_network_hub[each.value].resource_group_name
  resource_network_peer_role = "hub"
}

module "network_peering_hub" {
  for_each                   = var.service_location
  source                     = "../"
  service_environment        = var.service_environment
  resource_network_peer      = var.service_network_hub[each.value].network_name
  resource_group_peer        = var.service_network_hub[each.value].resource_group_name
  resource_network_peer_role = "spoke"
}