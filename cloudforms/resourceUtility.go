package cloudforms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/tidwall/gjson"
)

// getOrder : Function to check whether orders are present or not in order list
func getOrder(config Config, d *schema.ResourceData) error {
	url := "api/service_orders?expand=resources"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("[ERROR] Error in creating http Request %s", err)
		return err
	}
	response, err := config.GetResponse(request)
	if err != nil {
		log.Printf("[ERROR] Error in getting response %s", err)
		return err
	}
	data := string(response)
	subCountFromResult := gjson.Get(data, "subcount")
	subCount := subCountFromResult.Uint() //convert json result type to int
	if subCount == 0 {
		fmt.Println("Service order not found")
		log.Println("[ERROR] Service order not found")
		d.SetId("")
	}
	return nil
}

// deleteOrder : Function to delete an order with orderID corresponds to given requestID
func deleteOrder(config Config, orderID string) error {

	url := "api/service_orders/" + orderID

	// buff will contain request body
	buff, err := json.Marshal(map[string]string{
		"action": "delete",
	})
	if err != nil {
		log.Fatalln(err)
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(buff))
	if err != nil {
		log.Printf("[ERROR] Error in creating http Request %s", err)
	}

	response, err := config.GetResponse(request)
	if err != nil {
		log.Printf("[ERROR] Error in getting response %s", err)
	}

	log.Println(string(response))
	return nil

}

// checkrequestStatus : Function to check request status
func checkrequestStatus(d *schema.ResourceData, config Config, requestID string, timeOut int) error {
	timeout := time.After(time.Duration(timeOut) * time.Second)
	for {
		select {
		case <-time.After(1 * time.Second):
			status, state, err := checkServiceRequestStatus(config, requestID)
			if err == nil {
				if state == "finished" && status == "Ok" {
					log.Println("[DEBUG] Service order added SUCCESSFULLY")
					d.SetId(requestID)
					return nil
				} else if status == "Error" {
					log.Println("[ERROR] Failed")
					return fmt.Errorf("[Error] Failed execution")
				} else {
					log.Println("[DEBUG] Request state is :", state)
				}
			} else {
				return err
			}
		case <-timeout:
			log.Println("[DEBUG] Timeout occured")
			return fmt.Errorf("[ERROR] Timeout")
		}
	}
}

// checkServiceRequestStatus : Function to fetch service request status
func checkServiceRequestStatus(config Config, requestID string) (string, string, error) {

	url := "api/service_requests/" + requestID
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("[ERROR] Error in creating http Request %s", err)
		return "", "", fmt.Errorf("[ERROR] Error in creating http Request %s", err)
	}
	response, err := config.GetResponse(request)
	if err != nil {
		log.Printf("[ERROR] Error in getting response %s", err)
		return "", "", fmt.Errorf("[ERROR] Error in getting response %s", err)
	}

	data := string(response)
	// get request status
	statusFromResult := gjson.Get(data, "status")
	status := statusFromResult.String()

	// get request_state
	requestStateFromResult := gjson.Get(data, "request_state")
	requestState := requestStateFromResult.String()

	return status, requestState, nil
}

// ReadJSON : Function to read json data from file and add href into it
func ReadJSON(inputFileName string, href string) ([]byte, []byte, error) {

	jsonFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	defer jsonFile.Close()
	fileData, _ := ioutil.ReadFile(inputFileName)
	var result map[string]map[string]interface{}
	var result1 map[string]interface{}
	json.Unmarshal([]byte(fileData), &result1) // for "action":
	json.Unmarshal([]byte(fileData), &result)  //  for "resource" :

	// will add key value into map
	result["resource"]["href"] = href

	buff1, _ := json.Marshal(&result1)
	buff2, _ := json.Marshal(&result)

	return buff1, buff2, nil
}
