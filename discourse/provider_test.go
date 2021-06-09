package discourse

import(
	"os"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"testing"
	"log"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	os.Setenv("API_USERNAME","Ashwinigaddagiwork")
	os.Setenv("API_KEY","7c27290e2ccca7ae4427dfe518f68fb0659cd7df9908ff1cefbf034f9900a568")
	os.Setenv("BASE_URL", "https://clevertaptest.trydiscourse.com")
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"discourse": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		log.Println("[ERROR]: ",err)
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T)  {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("API_USERNAME"); v == "" {
		t.Fatal("API_USERNAME must be set for acceptance tests")
	}
	if v := os.Getenv("API_KEY"); v == "" {
		t.Fatal("API_KEY must be set for acceptance tests")
	}
	if v := os.Getenv("BASE_URL"); v == "" {
		t.Fatal("BASE_URL must be set for acceptance tests")
	}
}
