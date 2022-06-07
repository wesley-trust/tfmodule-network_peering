package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestValidateNetworkPeering(t *testing.T) {

	// Generate a random ID to prevent a naming conflict
	uniqueID := random.UniqueId()
	testREF := "NetworkPeering"
	serviceDeployment := testREF + "-" + uniqueID

	// Define variables
	locations := []string{"UK South"}

	// Validate module
	// Enable retryable error
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		// The path to where the Terraform code is located
		TerraformDir: "../examples/",

		// Variables to pass to the Terraform code using -var options
		Vars: map[string]interface{}{
			"service_deployment": serviceDeployment,
			"service_location":   locations,
		},
	})

	// Run `terraform init` and `terraform validate`. Fail the test if there are any errors.
	terraform.InitAndValidate(t, terraformOptions)
}

/* func TestPlanNetworkPeering_ApplyDependencies(t *testing.T) {

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

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.DestroyE(t, terraformDependencyOptions)

	// At the end of the test, run `terraform destroy` again, as replication delays can cause the initial destroy to fail
	defer terraform.Destroy(t, terraformDependencyOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformDependencyOptions)

	// Plan module
	// Enable retryable error
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		// The path to where the Terraform code is located
		TerraformDir: "../examples/",

		// Variables to pass to the Terraform code using -var options
		Vars: map[string]interface{}{
			"service_deployment": uniqueID,
			"service_location":   locations,
		},
	})

	// Run `terraform init` and `terraform plan`. Fail the test if there are any errors.
	terraform.InitAndPlan(t, terraformOptions)
} */
