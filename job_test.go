package coconut

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const INPUT_URL = "https://s3-eu-west-1.amazonaws.com/files.coconut.co/bbb_800k.mp4"

func TestRetrieveJob(t *testing.T) {
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

	j_ := cli.Job.Retrieve(j.Id)

	assert.Equal(t, j_.Id, j.Id)
}

func TestCreateJob(t *testing.T) {
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

	j, err := cli.Job.Create(JobCreate{
		Input: InputCreate{
			"url": INPUT_URL,
		},
		Outputs: OutputCreate{
			"mp4:240p": OutputParams{
				"path": "/video.mp4",
			},
		},
	})

	assert.Nil(t, err)
	assert.NotNil(t, j.Id)
}

func TestCreateJobWithError(t *testing.T) {
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

	_, err := cli.Job.Create(JobCreate{
		Input: InputCreate{
			"url": "notvalidurl",
		},
		Outputs: OutputCreate{
			"mp4:240p": OutputParams{
				"path": "/video.mp4",
			},
		},
	})

	assert.NotNil(t, err)
	assert.Equal(t, "input_url_not_valid", err.(Error).Code)
}
