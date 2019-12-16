package cloudforms

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

type res struct {
	Href              string      `json:"href"`
	ID                string      `json:"id"`
	Description       string      `json:"description"`
	ApprovalState     string      `json:"approval_state"`
	Type              string      `json:"type"`
	CreatedOn         time.Time   `json:"created_on"`
	UpdatedOn         time.Time   `json:"updated_on"`
	FulfilledOn       interface{} `json:"fulfilled_on"`
	RequesterID       string      `json:"requester_id"`
	RequesterName     string      `json:"requester_name"`
	RequestType       string      `json:"request_type"`
	RequestState      string      `json:"request_state"`
	Message           string      `json:"message"`
	Status            string      `json:"status"`
	Options           Options     `json:"options"`
	Userid            string      `json:"userid"`
	SourceID          string      `json:"source_id"`
	SourceType        string      `json:"source_type"`
	DestinationID     interface{} `json:"destination_id"`
	DestinationType   interface{} `json:"destination_type"`
	TenantID          string      `json:"tenant_id"`
	ServiceOrderID    string      `json:"service_order_id"`
	Process           bool        `json:"process"`
	CancelationStatus interface{} `json:"cancelation_status"`
	Actions           []Actions   `json:"actions"`
}

func TestAccServiceRequest_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckServiceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckServiceConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckServiceExists("cloudforms_service_request.test"),
				),
			},
		},
	})
}
func testAccCheckServiceDestroy(s *terraform.State) error {
	return nil
}
func testAccCheckServiceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No request ID is set")
		}
		config := testAccProvider.Meta().(Config)
		url := "api/service_requests/" + rs.Primary.ID
		req, err := http.NewRequest("GET", url, nil)
		resp, err := config.GetResponse(req)
		if err != nil {
			log.Fatal(err)
		}
		var record res
		if err := json.Unmarshal(resp, &record); err != nil {
			log.Println(err)
		}
		if record.Status != "Ok" {
			return fmt.Errorf("[ERROR] Service is not executed")
		}
		return nil
	}
}
func testAccCheckServiceConfig() string {
	return fmt.Sprintf(`
	provider "cloudforms" {
	ip       = "%s"  
	user_name  = "%s"
	password = "%s"
	  
	}

	# Data Source cloudforms_service_template
	data "cloudforms_service_template" "mytemplate"{
		name = "%s"
	}
	

	# Resource cloudforms_service_request
	resource "cloudforms_service_request" "test" {  
		name = "%s"
		template_href = "${data.cloudforms_service_template.mytemplate.href}"
		catalog_id ="${data.cloudforms_service_template.mytemplate.service_template_catalog_id}"
		input_file_name = "%s"
		time_out= 50
	} 
	
	`,
		os.Getenv("CF_SERVER_IP"),
		os.Getenv("CF_USER_NAME"),
		os.Getenv("CF_PASSWORD"),
		os.Getenv("CF_TEMPLATE_NAME"),
		os.Getenv("CF_TEMPLATE_NAME"),
		os.Getenv("CF_INPUT_FILE_NAME"))
}
