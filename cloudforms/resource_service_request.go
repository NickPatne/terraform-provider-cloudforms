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
func resourceServiceRequestCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(Config)
	inputFileName := d.Get("input_file_name").(string)
	timeout := d.Get("time_out").(int)
	href := d.Get("template_href").(string)
	catalogID := d.Get("catalog_id").(string)

	// templateStruct : struct to store attributes of service
	var templateStruct template

	url := "api/service_catalogs/" + catalogID + "/service_templates"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("[ERROR] Error in creating http Request %s", err)
	}
	response, err := config.GetResponse(request)
	if err != nil {
		log.Printf("[ERROR] Error in getting response %s", err)
	}

	// caling helper function to write href into file
	file, file1 := ReadJSON(inputFileName, href)

	err = json.Unmarshal(file, &templateStruct)
	if err != nil {
		log.Printf("[ERROR] Error while unmarshal file's json %s", err)
		return fmt.Errorf("[ERROR] Error while unmarshal file's json %s", err)
	}

	err = json.Unmarshal(file1, &templateStruct)
	if err != nil {
		log.Printf("[ERROR] Error while unmarshal file's json %s", err)
		return fmt.Errorf("[ERROR] Error while unmarshal file's json %s", err)
	}

	buff, _ := json.Marshal(&templateStruct)
	log.Println("[DEBUG] template struct:", templateStruct)

	var jsonStr = []byte(buff)
	log.Println("[DEBUG] json str in string form :", string(jsonStr))
	url = "api/service_catalogs/" + catalogID + "/service_templates"
	request, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Printf("[ERROR] Error in creating http Request %s", err)
		return fmt.Errorf("[ERROR] Error in creating http Request %s", err)
	}
	response, err = config.GetResponse(request)
	if err != nil {
		log.Printf("[ERROR] Error in getting response %s", err)
		return fmt.Errorf("[ERROR] Error in getting response %s", err)
	}

	// requestGostruct : struct to store response of post request
	var requestGostruct requestJsonstruct
	if err = json.Unmarshal(response, &requestGostruct); err != nil {
		log.Printf("[ERROR] Error while unmarshal requests json %s", err)
		return fmt.Errorf("[ERROR] Error while unmarshal requests json %s", err)
	}

	requestID := requestGostruct.Results[0].ID
	log.Println("[DEBUG] request id:", requestID)

	// check for timeout
	if timeout == 0 {
		return checkrequestStatus(d, config, requestID, 50)
	} else {
		return checkrequestStatus(d, config, requestID, timeout)
	}
}

// resourceServiceRequestRead : This function will read resource
func resourceServiceRequestRead(d *schema.ResourceData, m interface{}) error {
	config := m.(Config)
	err := getOrder(config, d)
	return err

}

// resourceServiceRequestDelete : This function will delete resource
func resourceServiceRequestDelete(d *schema.ResourceData, m interface{}) error {

	resourceServiceRequestRead(d, m)
	if d.Id() == "" {
		log.Println("[ERROR] Cannot find Order")
		return fmt.Errorf("[ERROR] Cannot find Order")
	}
	config := m.(Config)

	url1 := "api/service_requests/" + d.Id()
	req, err := http.NewRequest("GET", url1, nil)
	if err != nil {
		log.Printf("[ERROR] Error in creating http Request %s", err)
		return err
	}
	response, err := config.GetResponse(req)

	data3 := string(response)
	oi := gjson.Get(data3, "service_order_id")
	oID := oi.String()
	err1 := deleteOrder(config, oID)
	return err1
}
