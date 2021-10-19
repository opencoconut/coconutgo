# Coconut Go Library

The Coconut Go library provides access to the Coconut API for encoding videos, packaging media files into HLS and MPEG-Dash, generating thumbnails and GIF animation.

This library is only compatible with the Coconut API v2.

## Documentation

See the [full documentation](https://docs.coconut.co).

## Installation

```console
go get github.com/opencoconut/coconutgo
```

## Usage

The library needs you to set your API key which can be found in your [dashboard](https://app.coconut.co/api). Webhook URL and storage settings are optional but are very convenient because you set them only once.

### Example

```go
package main

import (
  "fmt"
  "os"
  "github.com/opencoconut/coconutgo"
)

func main() {
  // Initialize the Client
  cli := coconut.NewClient(coconut.Client{
    APIKey: os.Getenv("COCONUT_API_KEY"),
    Storage: coconut.Storage{
      "service": "s3",
      "region":  "us-east-1",
      "credentials": coconut.StorageCredentials{
        "access_key_id":     os.Getenv("AWS_ACCESS_KEY_ID"),
        "secret_access_key": os.Getenv("AWS_SECRET_ACCESS_KEY"),
      },
      "bucket": "mybucket",
    },
    Notification: coconut.Notification{
      "type": "http",
      "url":  "https://yoursite/api/coconut/webhook",
    },
  })

  // Create a job
  job, err := cli.Job.Create(coconut.JobCreate{
    Input: coconut.InputCreate{
      "url": "https://mysite/path/file.mp4",
    },
    Outputs: coconut.OutputCreate{
      "jpg:300x": coconut.OutputParams{
        "path": "/image.jpg",
      },
      "mp4:1080p": coconut.OutputParams{
        "path": "/1080p.mp4",
      },
      "httpstream": coconut.OutputParams{
        "hls": coconut.OutputParams{
          "path": "hls/",
        },
      },
    },
  })

  if err != nil {
    fmt.Printf("%# v", err)
  } else {
    fmt.Printf("%# v", job)
  }
}
```

### Choose the region

```go
cli.Region = "eu-west-1"
```

### Enabling Ultrafast Mode

```go
job, err := cli.Job.Create(coconut.JobCreate{
  Settings: coconut.Settings{
    "ultrafast": true
  },
  Input: coconut.InputCreate{
    "url": "https://mysite/path/file.mp4",
  },
  Outputs: coconut.OutputCreate{
    "mp4:2160p": coconut.OutputParams{
      "path": "/4k.mp4",
    },
  },
})
```

## Getting information about a job

```go
job := cli.Job.Retrieve("TsySPignC2xhOK")

for i, o := range job.Outputs {
  if o.Type == "video" {
    fmt.Printf("%d) Video: %# v\n\n", i, o.GetVideoURL())
  } else if o.Type == "image" {
    fmt.Printf("%d) Image: %# v\n\n", i, o.GetImageURLs())
  } else if o.Type == "httpstream" {
    fmt.Printf("%d) HTTPStream: %# v\n\n", i, o.GetHTTPStreamURLs())
  }
}
```

## Retrieving metadata

```go
cli.Metadata.retrieve("OolQXaiU86NFki")
```

*Released under the [MIT license](http://www.opensource.org/licenses/mit-license.php).*