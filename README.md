# Go client Library for encoding Videos with Coconut

## Install

```console
go get github.com/opencoconut/coconutgo
```

## Submitting the job

Use the [API Request Builder](https://app.coconut.co/jobs/new) to generate a config file that match your specific workflow.

Example of `coconut.conf`:

```ini
var s3 = s3://accesskey:secretkey@mybucket

set webhook = http://mysite.com/webhook/coconut

-> mp4  = $s3/videos/video.mp4
-> webm = $s3/videos/video.webm
-> jpg:300x = $s3/previews/thumbs_#num#.jpg, number=3
```

Here is the Go code to submit the config file:

```go
package main

import (
  "fmt"
  "github.com/opencoconut/coconutgo"
)

func main() {
  config := coconut.Config{
    Conf: "coconut.conf",
    Source: "http://yoursite.com/media/video.mp4",
    Vars: coconut.Vars{"vid": "1234"},
  }

  if job, err := coconut.NewJob(config, "api-key"); err != nil {
    fmt.Println("Error:", err)
  } else {
    fmt.Println("Job created:", job.Id)
  }
}
```

You can also create a job without a config file. To do that you will need to give every settings in the method parameters. Here is the exact same job but without a config file:

```go
vid := "1234"
s3 := "s3://accesskey:secretkey@mybucket"

config := coconut.Config{
  Vars: coconut.Vars{
    "vid": vid,
    "s3": s3,
  },
  Source: "http://yoursite.com/media/video.mp4",
  Webhook: "http://mysite.com/webhook/coconut?videoId=$vid",
  Outputs: coconut.Outputs{
    "mp4": "$s3/videos/video_$vid.mp4",
    "webm": "$s3/videos/video_$vid.webm",
    "jpg:300x": "$s3/previews/thumbs_#num#.jpg, number=3",
  }
}

if job, err := coconut.NewJob(config, "api-key"); err != nil {
  fmt.Println("Error:", err)
} else {
  fmt.Println("Job created:", job.Id)
}
```

Note that you can use the environment variable `COCONUT_API_KEY` to set your API key.


## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Added some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

*Released under the [MIT license](http://www.opensource.org/licenses/mit-license.php).*

---

* Coconut website: http://coconut.co
* API documentation: http://coconut.co/docs
* Contact: [support@coconut.co](mailto:support@coconut.co)
* Twitter: [@OpenCoconut](http://twitter.com/opencoconut)
