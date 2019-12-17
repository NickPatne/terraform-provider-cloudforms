package cloudforms

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

// tokenJsonstruct : contains token attributes
type tokenJsonstruct struct {
	AuthToken string `json:"auth_token"`
	TokenTTL  int    `json:"token_ttl"`
	ExpiresOn string `json:"expires_on"`
}

// GetToken : This funtion will generate Access Token
func GetToken(IP, UserID, Password string) (string, error) {
	token, err := getTokenFromServer(IP, UserID, Password)
	if err != nil {
		log.Printf("[Error] Cannot get Token : %s", err.Error())
		return "", err
	}
	return token, nil
}

// getTokenFromServer : helper funtion to generate access token
func getTokenFromServer(IP, UserID, Password string) (string, error) {
	req, err := http.NewRequest("GET", "https:/"+"/"+IP+"/api/auth", nil)
	if err != nil {
		log.Println("[ERROR] Error while requesting Token", err)
		return "", err
	}
	transportFlag := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// set request header
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(UserID, Password)

	client := &http.Client{Transport: transportFlag}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("[ERROR] Error while requesting Token", err)
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[Error] Error while getting response: %s", err)
	}

	// set request header
	if resp.StatusCode == http.StatusOK {
		var tokenStruct tokenJsonstruct
		if err = json.Unmarshal(body, &tokenStruct); err != nil {
			log.Printf("[ERROR] Error while unmarshal %s", err)
			return "", fmt.Errorf("[ERROR] Error while unmarshal %s", err)
		}
		token := tokenStruct.AuthToken
		return token, nil
	}
	return "", fmt.Errorf(httpResponseStatus(resp))
}

// GetGroup : will return User's Group name from MIQ-Server
func GetGroup(IP, UserID, Password string) (string, error) {
	group, err := getUserGroup(IP, UserID, Password)
	if err != nil {
		log.Printf("[Error] Cannot get Token : %s", err.Error())
		return "", err
	}
	return group, nil
}

// getUserGroup : helper funtion to get user group
func getUserGroup(IP, UserID, Password string) (string, error) {

	req, err := http.NewRequest("GET", "https:/"+"/"+IP+"/api", nil)
	if err != nil {
		log.Println("[ERROR] Error while requesting Token", err)
		return "", err
	}
	transportFlag := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(UserID, Password)

	client := &http.Client{Transport: transportFlag}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("[ERROR Error while requesting Token]", err)
		return "", err
	}
	if resp.StatusCode == http.StatusOK {

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("[Error]")
		}

		data := string(body)
		// used gjson package to get required value from json
		group := gjson.Get(data, "identity.group")

		return group.String(), nil
	}
	return "", fmt.Errorf(httpResponseStatus(resp))
}
