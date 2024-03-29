# Define variables
variable "service_name" {
  description = "Desired name for the provisioned resources"
  type        = string
  default     = "Dependency"
}

variable "service_environment" {
  description = "Desired environment for the service collection of provisioned resources"
  type        = string
  default     = "Test"
}

variable "service_location" {
  description = "The production resource locations to deploy"
  type        = list(string)
  default = [
    "UK South",
    "North Central US"
  ]
}

variable "service_deployment" {
  description = "Desired deployment identifier of the service collection of provisioned resources"
  type        = string
}

variable "resource_address_space" {
  description = "Desired address space for the provisioned resources"
  type        = map(any)
  default = {
    "Spoke" = {
      "UK South"         = "10.0.2.0/24"
      "North Central US" = "10.6.2.0/24"
    }
    "Hub" = {
      "UK South"         = "10.0.0.0/24"
      "North Central US" = "10.6.0.0/24"
    }
  }
}
