package coconut

import (
	"encoding/json"
	"fmt"
)

type Job struct {
	Id          string                   `json:"id"`
	Status      string                   `json:"status"`
	CreatedAt   string                   `json:"created_at"`
	CompletedAt string                   `json:"completed_at"`
	Progress    string                   `json:"progress"`
	Input       Input                    `json:"input"`
	Outputs_    []map[string]interface{} `json:"outputs"`
	Outputs     []Output
}

type Input struct {
	Status string
}

type Output struct {
	Status string
	Key    string
	Type   string
	Format string
	url    interface{}
}

type JobCreate struct {
	Input        InputCreate
	Storage      Storage
	Notification Notification
	Outputs      OutputCreate
}

type InputCreate map[string]interface{}

type OutputCreate map[string]interface{}
type OutputParams map[string]interface{}

func (d JobCreate) ToParams() P {
	m := make(map[string]interface{})
	m["input"] = d.Input
	m["storage"] = d.Storage
	m["notification"] = d.Notification
	m["outputs"] = d.Outputs

	return m
}

func (o Output) GetVideoURL() string {
	if o.Type != "video" {
		return ""
	}

	url := fmt.Sprintf("%v", o.url)

	return url
}

func (o Output) GetImageURLs() []string {
	if o.url == nil || o.Type != "image" {
		return make([]string, 0)
	}

	urls := make([]string, 0)
	for _, u := range o.url.([]interface{}) {
		urls = append(urls, fmt.Sprintf("%v", u))
	}

	return urls
}

func (o Output) GetHTTPStreamURLs() []map[string]string {
	if o.url == nil || o.Type != "httpstream" {
		return make([]map[string]string, 0)
	}

	urls := make([]map[string]string, 0)
	for _, u := range o.url.([]interface{}) {
		urls = append(urls, map[string]string{
			"format": fmt.Sprintf("%v", u.(map[string]interface{})["format"]),
			"url":    fmt.Sprintf("%v", u.(map[string]interface{})["url"]),
		})
	}

	return urls
}

func NewJob(dict map[string]interface{}) Job {
	jsonbody, err := json.Marshal(dict)
	if err != nil {
		return Job{}
	}

	j := Job{}
	if err := json.Unmarshal(jsonbody, &j); err != nil {
		return Job{}
	}

	for _, o := range j.Outputs_ {
		output := Output{
			Status: fmt.Sprintf("%v", o["status"]),
			Key:    fmt.Sprintf("%v", o["key"]),
			Type:   fmt.Sprintf("%v", o["type"]),
			Format: fmt.Sprintf("%v", o["format"]),
		}

		if o["type"] == "video" {
			url := o["url"]
			if url == nil {
				url = ""
			}
			output.url = url

		} else {
			output.url = o["urls"]
		}

		j.Outputs = append(j.Outputs, output)
	}

	return j
}

func (c *JobClient) Retrieve(jid string) Job {
	api := API{client: c.Client}
	res, err := api.Request("GET", "/jobs/"+jid, nil)

	if err != nil {
		return Job{}
	}

	return NewJob(res)
}

func (c *JobClient) Create(data JobCreate) (Job, error) {
	api := API{client: c.Client}

	if c.Client.Notification != nil {
		data.Notification = c.Client.Notification
	}

	if c.Client.Storage != nil {
		data.Storage = c.Client.Storage
	}

	res, err := api.Request("POST", "/jobs", data.ToParams())

	if err != nil {
		return Job{}, err
	}

	return NewJob(res), nil
}
