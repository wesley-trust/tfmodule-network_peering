# Required service variables
variable "service_environment" {
  description = "Desired environment for the service collection of provisioned resources"
  type        = string
}

# Required resource variables
variable "resource_network_peer" {
  description = "Resource outputs for peering"
  type = map(any)
}

variable "resource_network_peer_role" {
  description = "The peering type"
  type        = string
}