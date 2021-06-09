package discourse

import(
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
	"os"
)

func TestAccUser_Basic(t *testing.T) {
	os.Setenv("TF_ACC", "1")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckUserBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("discourse_user.user1", "email", "ashutosh.verma@clevertap.com"),
				),
			},
		},
	})
}

func testAccCheckUserBasic() string {
	return fmt.Sprintf(`
		resource "discourse_user" "user1" {
			email = "ashutosh.verma@clevertap.com"
		}
	`)
}

func TestAccUser_Update(t *testing.T) {
	os.Setenv("TF_ACC", "1")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckUserUpdatePre(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"discourse_user.user1", "email", "ashutosh.verma@clevertap.com"),
					resource.TestCheckResourceAttr(
						"discourse_user.user1", "name", "Ashutosh Verma"),	
				),
			},
			{
				Config: testAccCheckUserUpdatePost(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"discourse_user.user1", "email", "ashutosh.verma@clevertap.com"),
					resource.TestCheckResourceAttr(
						"discourse_user.user1", "name", "Ashutosh Testing"),
				),
			},
		},
	})
}

func testAccCheckUserUpdatePre() string {
	return fmt.Sprintf(`
		resource "discourse_user" "user1" {
			email = "ashutosh.verma@clevertap.com"
  			name = "Ashutosh Verma"
		}
	`)
}

func testAccCheckUserUpdatePost() string {
	return fmt.Sprintf(`
		resource "discourse_user" "user1" {
			email = "ashutosh.verma@clevertap.com"
  			name = "Ashutosh Testing"
		}
	`)
}
