module "network_peering_hub" {
  depends_on = [
    module.service_network_hub
  ]
  for_each                   = toset(var.service_location)
  source                     = "../"
  service_environment        = var.service_environment
  resource_network_peer      = module.service_network_hub[each.value].network_name
  resource_group_peer        = module.service_network_hub[each.value].resource_group_name
  resource_network_peer_role = "hub"
}

module "network_peering_spoke" {
    depends_on = [
    module.service_network_spoke
  ]
  for_each                   = toset(var.service_location)
  source                     = "../"
  service_environment        = var.service_environment
  resource_network_peer      = module.service_network_spoke[each.value].network_name
  resource_group_peer        = module.service_network_spoke[each.value].resource_group_name
  resource_network_peer_role = "spoke"
}