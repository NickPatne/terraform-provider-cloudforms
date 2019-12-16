package cloudforms

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccDataSourceServiceDetail_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckDataServiceDetailConfigServiceName(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudforms_service.myservice", "name", os.Getenv("CF_SERVICE_NAME")),
				),
			},
		},
	})
}

func testAccDataSourceServiceDetail(src, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		service := s.RootModule().Resources[src]
		serviceResource := service.Primary.Attributes

		search := s.RootModule().Resources[n]
		searchResource := search.Primary.Attributes

		testArrtributes := []string{
			"href",
			"id",
			"description",
			"tenant_id",
		}

		for _, attribute := range testArrtributes {
			if searchResource[attribute] != serviceResource[attribute] {
				return fmt.Errorf("Expected Service's parameter `%s` to be: `%s` but got: `%s`", attribute, serviceResource[attribute], searchResource[attribute])
			}
		}
		return nil
	}
}

func testAccCheckDataServiceDetailConfigServiceName() string {
	return fmt.Sprintf(`
	provider "cloudforms" {
		user_name  = "%s"
		password = "%s"
		ip       = "%s"
	  }

	data  "cloudforms_service" "myservice"{
		name = "%s"
	}
	`, os.Getenv("CF_USER_NAME"),
		os.Getenv("CF_PASSWORD"),
		os.Getenv("CF_SERVER_IP"),
		os.Getenv("CF_SERVICE_NAME"))
}
