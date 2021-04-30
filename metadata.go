package coconut

import (
	"encoding/json"
)

type Metadata struct {
	JobId    string                 `json:"job_id"`
	Metadata map[string]interface{} `json:"metadata"`
}

func NewMetadata(dict map[string]interface{}) Metadata {
	jsonbody, err := json.Marshal(dict)
	if err != nil {
		return Metadata{}
	}

	m := Metadata{}

	if err := json.Unmarshal(jsonbody, &m); err != nil {
		return Metadata{}
	}

	return m
}

func (c *MetadataClient) Retrieve(jid string) Metadata {
	api := API{client: c.Client}
	res, err := api.Request("GET", "/metadata/jobs/"+jid, nil)

	if err != nil {
		return Metadata{}
	}

	return NewMetadata(res)
}
