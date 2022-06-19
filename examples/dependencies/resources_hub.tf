module "resource_group_hub" {
  for_each            = toset(var.service_location)
  source              = "github.com/wesley-trust/tfmodule-resource_group"
  service_environment = var.service_environment
  service_name        = var.service_name
  service_location    = each.value
  service_deployment  = "${var.service_deployment}-Hub"
}

module "service_network_hub" {
  for_each                = toset(var.service_location)
  source                  = "github.com/wesley-trust/tfsubmodule-network"
  service_name            = "${var.service_name}-Hub"
  service_location_prefix = replace(each.value, "/[a-z[:space:]]/", "")
  resource_location       = module.resource_group_hub[each.value].location
  resource_group_name     = module.resource_group_hub[each.value].name
  resource_environment    = var.service_environment
  resource_deployment     = var.service_deployment
  resource_address_space  = lookup(lookup(var.resource_address_space, "Hub", null), each.value, null)
  resource_network_role   = "hub"
}
