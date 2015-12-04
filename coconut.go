package coconut

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ApiSettings struct {
	url       string
	endPoint  string
	userAgent string
}

type Error struct {
	Code    string `json:"error_code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s (code: %s)", e.Message, e.Code)
}

type Job struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

var api = ApiSettings{
	url:       "https://api.coconut.co",
	endPoint:  "/v1/job",
	userAgent: "Coconut/1.4.0 (Go)",
}

func NewJob(c Config, options ...string) (Job, error) {
	if conf, err := c.String(); err != nil {
		return Job{}, err
	} else {
		apiKey := ""
		// By default we get the API key from the env variable COCONUT_API_KEY
		if len(options) == 0 {
			apiKey = os.Getenv("COCONUT_API_KEY")
		} else {
			// API key is given in second parameter
			apiKey = options[0]
		}
		return Submit(conf, apiKey)
	}
}

func Submit(c string, apiKey string) (Job, error) {
	url := api.url + api.endPoint

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(c)))
	req.SetBasicAuth(apiKey, "")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", api.userAgent)
	req.Header.Set("Content-Type", "text/plain")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Job{}, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	// Job created successfully
	if resp.StatusCode == 201 {
		job := Job{}
		if err := json.Unmarshal([]byte(body), &job); err != nil {
			return Job{}, err
		} else {
			return job, nil
		}
	} else {
		coconutErr := Error{}
		if err := json.Unmarshal([]byte(body), &coconutErr); err != nil {
			return Job{}, err
		} else {
			return Job{}, coconutErr
		}
	}
}
