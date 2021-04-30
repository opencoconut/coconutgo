package coconut

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoconutClient(t *testing.T) {
	cli := Client{
		APIKey: "k-api-key",
	}

	c := cli
	assert.Equal(t, "k-api-key", c.APIKey)
	assert.Equal(t, "https://api.coconut.co/v2", c.GetEndpoint())
}

func TestCoconutClientWithRegion(t *testing.T) {
	cli := Client{
		APIKey: "k-api-key",
		Region: "us-west-2",
	}

	c := cli
	assert.Equal(t, "https://api-us-west-2.coconut.co/v2", c.GetEndpoint())
}

func TestCoconutClientWithEndpoint(t *testing.T) {
	cli := Client{
		APIKey:   "k-api-key",
		Endpoint: "http://localhost:3001/v2",
	}

	c := cli
	assert.Equal(t, "http://localhost:3001/v2", c.GetEndpoint())
}
