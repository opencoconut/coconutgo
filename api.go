package coconut

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const USER_AGENT = "Coconut/v2 GoBindings/2.0.0"

type API struct {
	client Client
}

type P map[string]interface{}

func (p P) String() string {
	jsonString, _ := json.Marshal(p)
	return string(jsonString)
}

func (a API) Request(m string, p string, data P) (map[string]interface{}, error) {
	url := a.client.GetEndpoint() + p

	var req *http.Request
	var err error

	if m == "GET" {
		req, err = http.NewRequest("GET", url, nil)
	} else if m == "POST" {
		req, err = http.NewRequest("POST", url, bytes.NewBuffer([]byte(data.String())))
	}

	req.SetBasicAuth(a.client.APIKey, "")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", USER_AGENT)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, Error{Code: "request_error", Message: "Request error", Status: "error"}
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// errors
	if resp.StatusCode > 399 {
		if resp.StatusCode == 400 || resp.StatusCode == 401 {
			var err = Error{}
			json.Unmarshal([]byte(body), &err)
			return nil, err
		} else {
			msg := fmt.Sprintf("Server returned HTTP status code %d", resp.StatusCode)
			return nil, Error{Code: "server_error", Message: msg, Status: "error"}
		}
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)

	return result, nil
}
