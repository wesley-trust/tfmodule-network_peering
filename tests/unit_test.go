package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestLocalNetworkPeering(t *testing.T) {
	t.Parallel()

	// Generate a random ID to prevent a naming conflict
	uniqueID := random.UniqueId()

	// Define variables
	locations := []string{"UK South"}

	// Plan dependencies
	// Enable retryable error
	terraformDependencyOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		// The path to where the Terraform code is located
		TerraformDir: "../examples/dependencies",

		// Variables to pass to the Terraform code using -var options
		Vars: map[string]interface{}{
			"service_deployment": uniqueID,
			"service_location":   locations,
		},
	})

	// Run `terraform init` and `terraform plan`. Fail the test if there are any errors.
	terraform.InitAndPlan(t, terraformDependencyOptions)

	// Define outputs
	serviceNetworkSpoke := terraform.OutputMap(t, terraformDependencyOptions, "service_network_spoke")
	serviceNetworkHub := terraform.OutputMap(t, terraformDependencyOptions, "service_network_hub")

	// Plan module
	// Enable retryable error
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		// The path to where the Terraform code is located
		TerraformDir: "../examples/",

		// Variables to pass to the Terraform code using -var options
		Vars: map[string]interface{}{
			"service_deployment":    uniqueID,
			"service_location":      locations,
			"service_network_spoke": serviceNetworkSpoke,
			"service_network_hub":   serviceNetworkHub,
		},
	})

	// Run `terraform init` and `terraform plan`. Fail the test if there are any errors.
	terraform.InitAndPlan(t, terraformOptions)
}
