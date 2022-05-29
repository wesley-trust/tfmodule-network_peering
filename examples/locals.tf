# Calculate local variables
locals {
  # Data queries
  service_network_hub   = data.terraform_remote_state.outputs.service_network_hub
  service_network_spoke = data.terraform_remote_state.outputs.service_network_spoke
}
