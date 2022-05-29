package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestValidateNetworkPeering(t *testing.T) {
	t.Parallel()

	// Enable retryable error
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		// The path to where the Terraform code is located
		TerraformDir: "../examples/",
	})

	// Run `terraform init` and `terraform validate`. Fail the test if there are any errors.
	terraform.InitAndValidate(t, terraformOptions)
}
