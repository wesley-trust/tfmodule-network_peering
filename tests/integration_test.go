package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestApplyNetworkPeering_Local(t *testing.T) {
	//t.Parallel()

	// Generate a random ID to prevent a naming conflict
	uniqueID := random.UniqueId()
	testREF := "NetworkPeering_Local"
	serviceDeployment := testREF + "-" + uniqueID

	// Define variables
	locations := []string{"UK South"}

	// Deploy dependencies
	// Enable retryable error
	terraformDependencyOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		// The path to where the Terraform code is located
		TerraformDir: "../examples/dependencies",

		// Variables to pass to the Terraform code using -var options
		Vars: map[string]interface{}{
			"service_deployment": serviceDeployment,
			"service_location":   locations,
		},
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.DestroyE(t, terraformDependencyOptions)

	// At the end of the test, run `terraform destroy` again, in case failures leave orphaned resources
	defer terraform.Destroy(t, terraformDependencyOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformDependencyOptions)

	// Deploy module
	// Enable retryable error
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		// The path to where the Terraform code is located
		TerraformDir: "../examples/",

		// Variables to pass to the Terraform code using -var options
		Vars: map[string]interface{}{
			"resource_network_peer_deployment": serviceDeployment,
			"service_location":                 locations,
		},
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.DestroyE(t, terraformOptions)

	// At the end of the test, run `terraform destroy` again, in case failures leave orphaned resources
	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)
}

func TestApplyNetworkPeering_Local_Global(t *testing.T) {
	//t.Parallel()

	// Generate a random ID to prevent a naming conflict
	uniqueID := random.UniqueId()
	testREF := "NetworkPeering_Local_Global"
	serviceDeployment := testREF + "-" + uniqueID

	// Define variables
	locations := []string{"UK South", "North Central US"}

	// Deploy dependencies
	// Enable retryable error
	terraformDependencyOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		// The path to where the Terraform code is located
		TerraformDir: "../examples/dependencies",

		// Variables to pass to the Terraform code using -var options
		Vars: map[string]interface{}{
			"service_deployment": serviceDeployment,
			"service_location":   locations,
		},
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.DestroyE(t, terraformDependencyOptions)

	// At the end of the test, run `terraform destroy` again, in case failures leave orphaned resources
	defer terraform.Destroy(t, terraformDependencyOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformDependencyOptions)

	// Deploy module
	// Enable retryable error
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		// The path to where the Terraform code is located
		TerraformDir: "../examples/",

		// Variables to pass to the Terraform code using -var options
		Vars: map[string]interface{}{
			"resource_network_peer_deployment": serviceDeployment,
			"service_location":                 locations,
		},
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.DestroyE(t, terraformOptions)

	// At the end of the test, run `terraform destroy` again, in case failures leave orphaned resources
	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)
}
