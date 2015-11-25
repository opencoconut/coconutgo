# Go client Library for encoding Videos with Coconut

## Install

```console
go get github.com/opencoconut/coconutgo
```

## Submitting the job

Use the [API Request Builder](https://app.coconut.co/job/new) to generate a config file that match your specific workflow.

Example of `heywatch.conf`:

```hw
var s3 = s3://accesskey:secretkey@mybucket

set source  = http://yoursite.com/media/video.mp4
set webhook = http://mysite.com/webhook/heywatch

-> mp4  = $s3/videos/video.mp4
-> webm = $s3/videos/video.webm
-> jpg_300x = $s3/previews/thumbs_#num#.jpg, number=3
```

Here is the Go code to submit the config file:

```go
package main

import (
  "fmt"
  "github.com/opencoconut/coconutgo"
  "io/ioutil"
)

func main() {
  config, _ := ioutil.ReadFile("coconut.conf")

  if job, err := coconut.Submit(string(config), "api-key"); err != nil {
    fmt.Println("Error:", err)
  } else {
    fmt.Println("Job created:", job["id"])
  }
}
```

*Released under the [MIT license](http://www.opensource.org/licenses/mit-license.php).*

---

* Coconut website: http://coconut.co
* API documentation: http://coconut.co/docs
* Contact: [support@coconut.co](mailto:support@coconut.co)
* Twitter: [@OpenCoconut](http://twitter.com/opencoconut)