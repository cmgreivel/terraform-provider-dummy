package dummy

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var testAccProvider *schema.Provider
var testAccProviders map[string]terraform.ResourceProvider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]terraform.ResourceProvider{
		"dummy": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	switch {
	case os.Getenv("DUMMY_PROVIDER_DIRECTORY") == "":
		t.Fatal("DUMMY_PROVIDER_DIRECTORY must be set for acceptance tests")
	}
}
