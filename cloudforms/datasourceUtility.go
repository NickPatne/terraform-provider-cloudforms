package cloudforms

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// GetServiceCatalog : Function will return service catalog
func GetServiceCatalog(config Config) ([]byte, error) {
	url := "api/service_catalogs"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("[ERROR] Error in creating request: %s ", err)
		return nil, err
	}

	expand := "resources"
	response, err := config.GetQueryResponse(request, expand, "", "")
	if err != nil {
		log.Printf("[ERROR] Error in getting response: %s ", err)
		return nil, err
	}
	return response, nil
}

// GetQueryResponse : Will return response with Query parameters
func (c *Config) GetQueryResponse(request *http.Request, expand string, attribute string, filter string) ([]byte, error) {

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

	// Initialize HTTPS client to skip SSL certificate verification
	transportFlag := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	var tempURL *url.URL
	tempURL, err = url.Parse("https:/" + "/" + c.IP + "/" + request.URL.Path)
	if err != nil {
		log.Println("[ERROR] URL is not in correct format")
		return nil, err
	}

	request.URL = tempURL

	//Setting Query Parameters
	if expand != "" {
		q := request.URL.Query()
		q.Add("expand", expand)
		request.URL.RawQuery = q.Encode()
	}

	if attribute != "" {
		q := request.URL.Query()
		q.Add("attributes", attribute)
		request.URL.RawQuery = q.Encode()
	}

	if filter != "" {
		q := request.URL.Query()
		q.Add("filter[]", filter)
		request.URL.RawQuery = q.Encode()
	}

	// Set headers for request
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "appliaction/json")

	// Set Custom headers for request
	request.Header.Set("X-Auth-Token", token)
	request.Header.Set("X-MIQ-Group", group)

	client := &http.Client{Transport: transportFlag}
	resp, err := client.Do(request)
	if err != nil {
		log.Println("[ERROR] Do: ", err)
		return nil, err
	}
	// check response StatusCode
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		log.Printf("[DEBUG] http Response StatusCode is %d with Text %s : ", resp.StatusCode, http.StatusText(resp.StatusCode))
		return ioutil.ReadAll(resp.Body)
	}
	return nil, fmt.Errorf(httpResponseStatus(resp))
}

// GetTemplateList : will return list of templates
func GetTemplateList(config Config, serviceID string) ([]byte, error) {
	url := "api/service_catalogs/" + serviceID
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("[ERROR] Error in creating request %s ", err)
		return nil, err
	}

	attribute := "service_templates"
	response, err := config.GetQueryResponse(request, "", attribute, "")
	if err != nil {
		log.Printf("[ERROR] Error in getting response %s ", err)
		return nil, err
	}
	return response, nil
}

// FlattenServiceTemplate : Helper Funtion to settle aggregate type of schema,
// Will store values within map interface
func FlattenServiceTemplate(list []ServiceTemplates) []map[string]interface{} {
	serviceTemplist := make([]map[string]interface{}, len(list))
	for i, serviceTemp := range list {
		l := map[string]interface{}{
			"href":         serviceTemp.Href,
			"id":           serviceTemp.ID,
			"name":         serviceTemp.Name,
			"description":  serviceTemp.Description,
			"guid":         serviceTemp.GUID,
			"miq_group_id": serviceTemp.MiqGroupID,
		}
		serviceTemplist[i] = l
	}
	return serviceTemplist
}

// httpResponseStatus : funtion to handle Client side errors
func httpResponseStatus(response *http.Response) string {
	var status string
	if response.StatusCode == http.StatusBadRequest {
		buffer, _ := ioutil.ReadAll(response.Body)
		log.Println(string(buffer))
		status = readErrorResponse(string(buffer))
	}
	if response.StatusCode == http.StatusNotFound {
		status = "[ Couldn't retrieve the content. 404 Not found. ]"
	}
	if response.StatusCode == http.StatusUnauthorized {
		status = "[ Invalid user login credentials, Please check username OR password ! ]"
	}
	return status
}

// GetServiceTemplate : Function to return template Details
func GetServiceTemplate(config Config, templateName string) ([]byte, error) {
	url := "api/service_templates" //GetTemplateAPI(templateName)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("[ERROR] Error in creating request %s ", err)
		return nil, err
	}
	expand := "resources"
	filter := "name=" + templateName
	response, err := config.GetQueryResponse(request, expand, "", filter)
	if err != nil {
		log.Printf("Response is %s ", string(response))
		log.Printf("[ERROR] Error in getting response [GetQueryResponse] \n %s ", err)
		return nil, err
	}
	log.Printf("Response is %s ", string(response))
	return response, nil
}

// readErrorResponse : Funtion to read error message
func readErrorResponse(data string) string {
	var errMsg ResponseError
	err := json.Unmarshal([]byte(data), &errMsg)

	if err != nil {
		fmt.Println(err)
	}
	log.Println("[>>>>DEBUG<<<<] Error Struct message ", errMsg.Error.Message)
	return errMsg.Error.Message
}
