package coconut

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestGenerateFullConfigWithNoFile(t *testing.T) {
	conf := Config{
		Vars: Vars{
			"vid":  "1234",
			"user": "5098",
			"s3":   "s3://a:s@bucket",
		},
		Source:  "https://s3-eu-west-1.amazonaws.com/files.coconut.co/test.mp4",
		Webhook: "http://mysite.com/webhook?vid=$vid&user=$user",
		Outputs: Outputs{
			"mp4":      "$s3/vid.mp4",
			"jpg_200x": "$s3/thumb.jpg",
			"webm":     "$s3/vid.webm",
		},
	}

	generated := strings.Join([]string{
		"var s3 = s3://a:s@bucket",
		"var user = 5098",
		"var vid = 1234",
		"",
		"set source = https://s3-eu-west-1.amazonaws.com/files.coconut.co/test.mp4",
		"set webhook = http://mysite.com/webhook?vid=$vid&user=$user",
		"",
		"-> jpg_200x = $s3/thumb.jpg",
		"-> mp4 = $s3/vid.mp4",
		"-> webm = $s3/vid.webm",
	}, "\n")

	c, _ := conf.String()
	assert.Equal(t, generated, c)
}

func TestGenerateConfigWithFile(t *testing.T) {
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

	generated := strings.Join([]string{
		"var s3 = s3://a:s@bucket",
		"var user = 5098",
		"var vid = 1234",
		"",
		"set source = https://s3-eu-west-1.amazonaws.com/files.coconut.co/test.mp4",
		"set webhook = http://mysite.com/webhook?vid=$vid&user=$user",
		"",
		"-> mp4 = $s3/vid.mp4",
	}, "\n")

	c, _ := conf.String()
	assert.Equal(t, generated, c)
}

func TestSetApiVersion(t *testing.T) {
	conf := Config{
		ApiVersion: "beta",
		Source:     "https://s3-eu-west-1.amazonaws.com/files.coconut.co/test.mp4",
		Webhook:    "http://mysite.com/webhook?vid=$vid&user=$user",
		Outputs: Outputs{
			"mp4": "$s3/vid.mp4",
		},
	}

	generated := strings.Join([]string{
		"set api_version = beta",
		"set source = https://s3-eu-west-1.amazonaws.com/files.coconut.co/test.mp4",
		"set webhook = http://mysite.com/webhook?vid=$vid&user=$user",
		"",
		"-> mp4 = $s3/vid.mp4",
	}, "\n")

	c, _ := conf.String()
	assert.Equal(t, generated, c)
}
