package cloudforms

import (
	"crypto/tls"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Config : Configuration required for provider connection
type Config struct {
	IP       string
	UserName string
	Password string
}

// CFConnect : will create client for connection
func CFConnect(d *schema.ResourceData) (interface{}, error) {

	ip := d.Get("ip").(string)
	// Check If field is not empty
	if ip == "" {
		return nil, fmt.Errorf("[ERROR] cloudforms server IP not found ")
	}

	username := d.Get("user_name").(string)
	// Check If field is not empty
	if username == "" {
		return nil, fmt.Errorf("[ERROR] cloudforms server username not found")
	}

	password := d.Get("password").(string)
	// Check If field is not empty
	if password == "" {
		return nil, fmt.Errorf("[ERROR] cloudforms server Password not found")
	}

	config := Config{
		IP:       ip,
		UserName: username,
		Password: password,
	}
	return config, nil
}

// GetResponse : Get the desired Response
// This function will return API response which will contain
// response.body in []byte format
func (c *Config) GetResponse(request *http.Request) ([]byte, error) {

	token, err := GetToken(c.IP, c.UserName, c.Password)
	if err != nil {
		log.Println("[ERROR] Error in getting token")
		return nil, err
	}
	// While authenticating with Token
	// it is necessary to provide user-group
	group, err := GetGroup(c.IP, c.UserName, c.Password)
	if err != nil {
		log.Println("[ERROR] Error in getting User group")
		return nil, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	var tempURL *url.URL
	tempURL, err = url.Parse("https:/" + "/" + c.IP + "/" + request.URL.Path)
	if err != nil {
		log.Println("[ERROR] URL is not in correct format")
		return nil, err
	}
	request.URL = tempURL
	log.Println(request.URL)

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "appliaction/json")
	request.Header.Set("X-Auth-Token", token)
	request.Header.Set("X-MIQ-Group", group)

	client := &http.Client{Transport: tr}
	resp, err := client.Do(request)
	if err != nil {
		log.Println("[ERROR] Error while getting response", err)
		return nil, err
	}
	// check response StatusCode
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		log.Printf("[DEBUG] http Response StatusCode is %d with Text %s : ", resp.StatusCode, http.StatusText(resp.StatusCode))
		return ioutil.ReadAll(resp.Body)
	}
	log.Printf("[DEBUG] http Response StatusCode is %d with Text %s : ", resp.StatusCode, http.StatusText(resp.StatusCode))
	return nil, fmt.Errorf(httpResponseStatus(resp))

}