package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func testVirtualGateway(t *testing.T, variant string) {
	t.Parallel()

	primaryHostedZone := os.Getenv("TF_VAR_primary_hosted_zone")

	if primaryHostedZone == "" {
		t.Fatal("TF_VAR_primary_hosted_zone must be set to run tests. e.g. 'export TF_VAR_primary_hosted_zone=example.org'")
	}

	terraformDir := fmt.Sprintf("../examples/%s", variant)

	terraformOptions := &terraform.Options{
		TerraformDir: terraformDir,
		LockTimeout:  "5m",
	}

	terraformTargetMeshOptions := &terraform.Options{
		TerraformDir: terraformDir,
		Targets: []string{
			"module.mesh",
		},
		LockTimeout: "5m",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformTargetMeshOptions)
	terraform.Apply(t, terraformOptions)

	name := terraform.Output(t, terraformOptions, "name")
	domainName := terraform.Output(t, terraformOptions, "domain_name")

	expectedName := fmt.Sprintf("example-tf-virtual-gateway-%s", variant)
	expectedDomainName := fmt.Sprintf("%s.%s", expectedName, primaryHostedZone)

	assert.Equal(t, expectedName, name)
	assert.Equal(t, expectedDomainName, domainName)
}
