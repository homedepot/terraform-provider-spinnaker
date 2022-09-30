package spinnaker

import (
	"os"
   "fmt"
	"testing"
   "context"

   "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
   "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"spinnaker": testAccProvider,
	}
}

func testAccPreCheck(t *testing.T) {
	if os.Getenv("GATE_URL") == "" {
		t.Fatal("GATE_URL must be set for acceptance tests")
	}
	diags := testAccProvider.Configure(context.Background(), terraform.NewResourceConfigRaw(nil))
	if diags != nil {
      err := ""
      for i, d := range diags {
         err = fmt.Sprintf("%s[%d]%s ", err, i, d.Summary)
      }
		t.Fatalf("err: %s", err)
	}
}

func TestProvider(t *testing.T) {
	//if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}
