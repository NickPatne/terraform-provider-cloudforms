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

//function to check weather orders are present or not in order list:
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
	sc := gjson.Get(data, "subcount")
	sc1 := sc.Uint() //convert json result type to int
	if sc1 == 0 {
		fmt.Println("No orders available")
		log.Println("[ERROR] Service order was not found")
		d.SetId("")
	}
	return nil
}

//Func to delete an order with oID corresponds to given requestID
func deleteOrder(config Config, oID string) error {

	url2 := "api/service_orders/" + oID
	reqBody, err := json.Marshal(map[string]string{
		"action": "delete",
	})
	if err != nil {
		log.Fatalln(err)
	}
	req, err := http.NewRequest("POST", url2, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Printf("[ERROR] Error in creating http Request %s", err)
	}

	response, err := config.GetResponse(req)
	if err != nil {
		log.Printf("[ERROR] Error in getting response %s", err)
	}

	log.Println(string(response))
	return nil

}

//Function to check request status:

func checkrequestStatus(d *schema.ResourceData, config Config, requestID string, timeOut int) error {
	timeout := time.After(time.Duration(timeOut) * time.Second)
	for {
		select {
		case <-time.After(1 * time.Second):
			status, state, err := getRequestResponse(config, requestID)
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

// function to fetch request response
func getRequestResponse(config Config, requestID string) (string, string, error) {

	url := "api/service_requests/" + requestID //service_request endpoint
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

	data2 := string(response)
	status1 := gjson.Get(data2, "status")
	status := status1.String()
	state1 := gjson.Get(data2, "request_state")
	state := state1.String()
	return status, state, nil
}

// ReadJSON : function to read json data from file and add href into it
func ReadJSON(inputFileName string, href string) ([]byte, []byte) {

	jsonFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadFile(inputFileName)
	var result map[string]map[string]interface{}
	var result1 map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result1) // for "action": "order",
	json.Unmarshal([]byte(byteValue), &result)  //  for "resource" : {}

	result["resource"]["href"] = href

	buff1, _ := json.Marshal(&result1)
	buff2, _ := json.Marshal(&result)

	return buff1, buff2

}
