package coconut

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const coconutURL = "https://api.coconut.co"

type CoconutError struct {
	Code    string `json:"error_code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func (e CoconutError) Error() string {
	return fmt.Sprintf("%s (code: %s)", e.Message, e.Code)
}

type CoconutJob struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

func Submit(Config string, APIKey string) (CoconutJob, error) {
	url := coconutURL + "/v1/job"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(Config)))
	req.SetBasicAuth(APIKey, "")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "HeyWatch/1.0.0 (Go)")
	req.Header.Set("Content-Type", "text/plain")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return CoconutJob{}, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	// Job created successfully
	if resp.StatusCode == 201 {
		job := CoconutJob{}
		if err := json.Unmarshal([]byte(body), &job); err != nil {
			return CoconutJob{}, err
		} else {
			return job, nil
		}
	} else {
		coconutErr := CoconutError{}
		if err := json.Unmarshal([]byte(body), &coconutErr); err != nil {
			return CoconutJob{}, err
		} else {
			return CoconutJob{}, coconutErr
		}
	}
}
