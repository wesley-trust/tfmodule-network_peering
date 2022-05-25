module "network_peering_spoke" {
  count                      = length(var.service_location)
  source                     = "../"
  service_environment        = var.service_environment
  resource_network_peer      = var.service_network_spoke
  resource_network_peer_role = "hub"
}

module "network_peering_hub" {
  count                      = length(var.service_location)
  source                     = "../"
  service_environment        = var.service_environment
  resource_network_peer      = var.service_network_hub
  resource_network_peer_role = "spoke"
}
