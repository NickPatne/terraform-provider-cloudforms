package cloudforms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/tidwall/gjson"
)

func resourceServiceRequest() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceRequestCreate,
		Read:   resourceServiceRequestRead,
		Delete: resourceServiceRequestDelete,

		Schema: map[string]*schema.Schema{
			// required values
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"input_file_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"template_href": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"catalog_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			// optional values
			"time_out": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

// resourceServiceRequestCreate : This function will create resource
func resourceServiceRequestCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)
	inputFileName := d.Get("input_file_name").(string)
	timeout := d.Get("time_out").(int)
	href := d.Get("template_href").(string)
	catalogID := d.Get("catalog_id").(string)

	// templateStruct : struct to store action and attributes of service
	var templateStruct template

	// calling helper function to write href into file
	file, file1, err := ReadJSON(inputFileName, href)
	if err != nil {
		log.Printf("[ERROR] %s", err)
		return fmt.Errorf("[ERROR] %s", err)
	}

	// will set action to struct
	err = json.Unmarshal(file, &templateStruct)
	if err != nil {
		log.Printf("[ERROR] Error while unmarshal file's json %s", err)
		return fmt.Errorf("[ERROR] Error while unmarshal file's json %s", err)
	}

	// will set resource attributes to struct
	err = json.Unmarshal(file1, &templateStruct)
	if err != nil {
		log.Printf("[ERROR] Error while unmarshal file's json %s", err)
		return fmt.Errorf("[ERROR] Error while unmarshal file's json %s", err)
	}

	// buff will contain request body
	buff, _ := json.Marshal(&templateStruct)

	url := "api/service_catalogs/" + catalogID + "/service_templates"
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(buff))
	if err != nil {
		log.Printf("[ERROR] Error in creating http Request %s", err)
		return fmt.Errorf("[ERROR] Error in creating http Request %s", err)
	}

	response, err := config.GetResponse(request)
	if err != nil {
		log.Printf("[ERROR] Error in getting response %s", err)
		return fmt.Errorf("[ERROR] Error in getting response %s", err)
	}

	// requestStruct : struct to store response body of post request
	var requestStruct requestJsonstruct
	if err = json.Unmarshal(response, &requestStruct); err != nil {
		log.Printf("[ERROR] Error while unmarshal requests json %s", err)
		return fmt.Errorf("[ERROR] Error while unmarshal requests json %s", err)
	}

	requestID := requestStruct.Results[0].ID
	log.Println("[DEBUG] request id:", requestID)

	// check for timeout
	if timeout == 0 {
		return checkrequestStatus(d, config, requestID, 180)
	} else {
		return checkrequestStatus(d, config, requestID, timeout)
	}
}

// resourceServiceRequestRead : This function will read resource
func resourceServiceRequestRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)
	err := getOrder(config, d)
	return err

}

// resourceServiceRequestDelete : This function will delete resource
func resourceServiceRequestDelete(d *schema.ResourceData, meta interface{}) error {

	// check whether resource exists
	resourceServiceRequestRead(d, meta)
	if d.Id() == "" {
		log.Println("[ERROR] Cannot find Order")
		return fmt.Errorf("[ERROR] Cannot find Order")
	}
	config := meta.(Config)

	url := "api/service_requests/" + d.Id()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("[ERROR] Error in creating http Request %s", err)
		return err
	}
	response, err := config.GetResponse(req)
	if err != nil {
		log.Printf("[ERROR] Error in getting response %s", err)
		return fmt.Errorf("[ERROR] Error in getting response %s", err)
	}

	data := string(response)
	// get service_order_id
	orderIDFromResult := gjson.Get(data, "service_order_id")
	orderID := orderIDFromResult.String()

	err = deleteOrder(config, orderID)
	return err
}
