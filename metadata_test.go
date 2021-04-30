package coconut

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRetrieveMetadata(t *testing.T) {
	cli := NewClient(Client{
		APIKey: os.Getenv("COCONUT_API_KEY"),
		Storage: Storage{
			"service": "s3",
			"region":  os.Getenv("AWS_REGION"),
			"credentials": StorageCredentials{
				"access_key_id":     os.Getenv("AWS_ACCESS_KEY_ID"),
				"secret_access_key": os.Getenv("AWS_SECRET_ACCESS_KEY"),
			},
			"bucket": os.Getenv("AWS_BUCKET"),
			"path":   "/coconutgo/tests/",
		},
		Notification: Notification{
			"type": "http",
			"url":  os.Getenv("COCONUT_WEBHOOK_URL"),
		},
	})

	j, _ := cli.Job.Create(JobCreate{
		Input: InputCreate{
			"url": INPUT_URL,
		},
		Outputs: OutputCreate{
			"mp4:240p": OutputParams{
				"path": "/video.mp4",
			},
		},
	})

	m := cli.Metadata.Retrieve(j.Id)

	assert.Equal(t, j.Id, m.JobId)
	assert.NotNil(t, m.Metadata)
}
