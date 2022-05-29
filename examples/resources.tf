module "network_peering_spoke" {
  count                      = length(var.service_location)
  source                     = "../"
  service_environment        = var.service_environment
  resource_network_peer      = var.service_network_spoke[count.index].network_name
  resource_group_peer        = var.service_network_spoke[count.index].resource_group_name
  resource_network_peer_role = "hub"
}

module "network_peering_hub" {
  count                      = length(var.service_location)
  source                     = "../"
  service_environment        = var.service_environment
  resource_network_peer      = var.service_network_hub[count.index].network_name
  resource_group_peer        = var.service_network_hub[count.index].resource_group_name
  resource_network_peer_role = "spoke"
}
