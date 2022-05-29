module "resource_group_spoke" {
  for_each            = toset(var.service_location)
  source              = "github.com/wesley-trust/tfmodule-resource_group"
  service_environment = var.service_environment
  service_name        = var.service_name
  service_location    = each.value
  service_deployment  = "${var.service_deployment}-Spoke"
}

module "service_network_spoke" {
  for_each                = toset(var.service_location)
  source                  = "github.com/wesley-trust/tfsubmodule-network"
  service_name            = "${var.service_name}-Spoke"
  service_location_prefix = replace(each.value, "/[a-z[:space:]]/", "")
  resource_location       = module.resource_group_spoke[each.value].location
  resource_group_name     = module.resource_group_spoke[each.value].name
  resource_environment    = var.service_environment
  resource_address_space  = lookup(lookup(var.resource_address_space, "Spoke", null), each.value, null)
  resource_network_role   = "spoke"
}