package cloudforms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/tidwall/gjson"
)

func resourceRequestMiq() *schema.Resource {
	return &schema.Resource{
		Create: resourceMiqOrder,
		Read:   resourceMiqGet,
		Delete: resourceMiqDelete,

		Schema: map[string]*schema.Schema{
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
			"time_out": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}
func resourceMiqOrder(d *schema.ResourceData, m interface{}) error {
	config := m.(Config)
	serviceName := d.Get("name").(string)
	inputFileName := d.Get("input_file_name").(string)
	timeout := d.Get("time_out").(int)
	response, err := getServiceCatalog(config)
	if err != nil {
		log.Printf("[ERROR] Error in getting response %s", err)
		return fmt.Errorf("[ERROR] Error in getting response %s", err)
	}
	var serviceCatalogGostruct ServiceCatalogJsonstruct

	if err = json.Unmarshal(response, &serviceCatalogGostruct); err != nil {
		log.Printf("[ERROR] Error while unmarshal requests json %s", err)
		return fmt.Errorf("[ERROR] Error while unmarshal requests json %s", err)
	}

	i := serviceCatalogGostruct.Subcount
	serviceID := ""
	for j := 0; j < i; j++ {
		if serviceCatalogGostruct.Resources[j].Name == serviceName {
			serviceID = serviceCatalogGostruct.Resources[j].ID
		}
	}
	if serviceID == "" {
		log.Printf("[Error] Service is not present")
		return fmt.Errorf("[ERROR] Service is not present %s", serviceName)
	}
	var t template
	url := "api/service_catalogs/" + serviceID + "/service_templates"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("[ERROR] Error in creating http Request %s", err)
	}
	response, err = config.GetResponse(request)
	if err != nil {
		log.Printf("[ERROR] Error in getting response %s", err)
	}
	file, e := ioutil.ReadFile(inputFileName)
	if e != nil {
		log.Printf("[ERROR] File error: %v\n", e)
	}

	err = json.Unmarshal(file, &t)
	if err != nil {
		log.Printf("[ERROR] Error while unmarshal file's json %s", err)
		return fmt.Errorf("[ERROR] Error while unmarshal file's json %s", err)
	}
	buff, _ := json.Marshal(&t)
	var jsonStr = []byte(buff)
	url = "api/service_catalogs/" + serviceID + "/service_templates"
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
	var requestGostruct requestJsonstruct
	if err = json.Unmarshal(response, &requestGostruct); err != nil {
		log.Printf("[ERROR] Error while unmarshal requests json %s", err)
		return fmt.Errorf("[ERROR] Error while unmarshal requests json %s", err)
	}
	requestID := requestGostruct.Results[0].ID
	log.Println("[DEBUG] request id:", requestID)
	if timeout == 0 {
		return checkrequestStatus(d, config, requestID, 50)
	} else {
		return checkrequestStatus(d, config, requestID, timeout)
	}
}

func resourceMiqGet(d *schema.ResourceData, m interface{}) error {
	config := m.(Config)
	err := getOrder(config, d)
	return err

}

func resourceMiqDelete(d *schema.ResourceData, m interface{}) error {

	resourceMiqGet(d, m)
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
