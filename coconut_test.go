package coconut

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

var apiKey = os.Getenv("COCONUT_API_KEY")

func TestSubmitJob(t *testing.T) {
	conf := Config{
		Source:  "https://s3-eu-west-1.amazonaws.com/files.coconut.co/test.mp4",
		Webhook: "http://mysite.com/webhook",
		Outputs: Outputs{"mp4": "s3://a:s@bucket/video.mp4", "webm": "s3://a:s@bucket/video.webm"},
	}

	if job, err := NewJob(conf); err != nil {
		t.Errorf("Error: %s", err)
	} else {
		t.Logf("Job created: %d", job.Id)
	}
}

func TestFailure(t *testing.T) {
	conf := Config{
		Source: "https://s3-eu-west-1.amazonaws.com/files.coconut.co/test.mp4",
	}

	if _, err := NewJob(conf); err != nil {
		assert.Equal(t, err.(Error).Status, "error")
		assert.Equal(t, err.(Error).Code, "config_not_valid")
	}
}

func TestSubmitJobWithAPIKey(t *testing.T) {
	conf := Config{
		Source:  "https://s3-eu-west-1.amazonaws.com/files.coconut.co/test.mp4",
		Webhook: "http://mysite.com/webhook",
		Outputs: Outputs{"mp4": "s3://a:s@bucket/video.mp4", "webm": "s3://a:s@bucket/video.webm"},
	}

	if _, err := NewJob(conf, "k-4d204a7fd1fc67fc00e87d3c326d9b75"); err != nil {
		assert.Equal(t, err.(Error).Status, "error")
		assert.Equal(t, err.(Error).Code, "authentication_failed")
	}
}

func TestSubmitFile(t *testing.T) {
	ioutil.WriteFile("coconut.conf", []byte("var s3 = s3://a:s@bucket\nset webhook = http://mysite.com/webhook?vid=$vid&user=$user\n-> mp4 = $s3/vid.mp4"), 0644)
	defer os.Remove("coconut.conf")

	conf := Config{
		Conf: "coconut.conf",
		Vars: Vars{
			"vid":  "1234",
			"user": "5098",
		},
		Source: "https://s3-eu-west-1.amazonaws.com/files.coconut.co/test.mp4",
	}

	if job, err := NewJob(conf); err != nil {
		t.Errorf("Error: %s", err)
	} else {
		t.Logf("Job created: %d", job.Id)
	}
}

func TestGetJob(t *testing.T) {
	conf := Config{
		Source:  "https://s3-eu-west-1.amazonaws.com/files.coconut.co/test.mp4",
		Webhook: "http://mysite.com/webhook",
		Outputs: Outputs{"mp4": "s3://a:s@bucket/video.mp4", "webm": "s3://a:s@bucket/video.webm"},
	}

	if job, err := NewJob(conf); err != nil {
		t.Errorf("Error: %s", err)
	} else {
		if info, err := GetJob(job.Id); err != nil {
			t.Errorf("Error: %s", err)
		} else {
			assert.Equal(t, info.Id, job.Id)
		}
	}
}
