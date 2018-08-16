# go-obs-websocket

[![Build Status](https://travis-ci.com/christopher-dG/go-obs-websocket.svg?branch=master)](https://travis-ci.com/christopher-dG/go-obs-websocket)
[![GoDoc](https://godoc.org/github.com/christopher-dG/go-obs-websocket?status.svg)](https://godoc.org/github.com/christopher-dG/go-obs-websocket)

`go-obs-websocket` provides client functionality for [`obs-websocket`](https://github.com/Palakis/obs-websocket).
Currently, the target version is `4.3`.

## Installation

```sh
go get github.com/christopher-dG/go-obs-websocket
```

## Usage

```go
package main

import (
	"log"
	"time"

	obs "github.com/christopher-dG/go-obs-websocket"
)

func main() {
	c := obs.Client{Host: "localhost", Port: 4444}
	if err := c.Connect(); err != nil {
		log.Fatal(err)
	}
	defer c.Disconnect()

	future, err := obs.NewGetStreamingStatusRequest().Send(c)
	if err != nil {
		log.Fatal(err)
	}

	status := (<-future)
	log.Println("streaming:", status.Streaming)

	c.AddEventHandler("Heartbeat", func(e obs.Event) {
		log.Println("profile:", e.(obs.HeartbeatEvent).CurrentProfile)
	})
	time.Sleep(time.Second * 10)
}
```

## gobs

```
Usage of ./gobs:
  -f string
    	JSON file to read requests from
  -host string
    	obs-websocket hostname (default "localhost")
  -password string
    	obs-websocket password
  -port int
    	obs-websocket port number (default 4444)
  -q	disable all output
```

This package also includes an executable: `gobs`.
It lets you send a series of requests from a list written in JSON, and the responses will be written to the console.

The only required field per entry is `request-type`.
You can see the list of requests and their fields [here](https://github.com/Palakis/obs-websocket/blob/4.3-maintenance/docs/generated/protocol.md).

There is also one supplemental request type: `Sleep`.
It takes a field named `seconds` and waits for that amount of time.

Here is an example JSON file:

```json
[
    {
        "request-type": "GetCurrentScene"
    },
    {
        "request-type": "Sleep",
        "seconds": 1
    },
    {
        "request-type": "SetRecordingFolder",
        "rec-folder": "/tmp"
    }
]
```
