# Required service variables
variable "service_environment" {
  description = "Desired environment for the service collection of provisioned resources"
  type        = string
}

variable "resource_network_peer_deployment" {
  description = "Deployment identifier for the resource network to peer"
  type        = string
}

# Required resource variables
variable "resource_network_peer" {
  description = "Resource network outputs for peering"
  type        = string
}

variable "resource_group_peer" {
  description = "Resource group outputs for peering"
  type        = string
}

variable "resource_network_peer_role" {
  description = "The peering type"
  type        = string
}
