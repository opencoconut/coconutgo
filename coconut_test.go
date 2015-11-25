package coconut

import (
	"os"
	"testing"
)

func TestSubmitJob(t *testing.T) {
	config := `
  set source = https://s3-eu-west-1.amazonaws.com/files.coconut.co/test.mp4
  set webhook = http://mysite.com/webhook
  -> mp4 = s3://a:s@bucket/video.mp4`

	if job, err := Submit(config, os.Getenv("COCONUT_API_KEY")); err != nil {
		t.Errorf("Error:", err)
	} else {
		t.Logf("Job created:", job.Id)
	}
}

func TestFailure(t *testing.T) {
	config := `
  set source = https://s3-eu-west-1.amazonaws.com/files.coconut.co/test.mp4
  `

	_, err := Submit(config, os.Getenv("COCONUT_API_KEY"))
	if err != nil {
		t.Logf("Error:", err)
	}
}
