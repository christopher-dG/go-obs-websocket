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
	// Connect a client.
	c := obs.Client{Host: "localhost", Port: 4444}
	if err := c.Connect(); err != nil {
		log.Fatal(err)
	}
	defer c.Disconnect()

	// Send and receive a request asynchronously.
	req := obs.NewGetStreamingStatusRequest()
	if err := req.Send(c); err != nil {
		log.Fatal(err)
	}
	// This will block until the response comes.
	resp, err := req.Receive()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("streaming:", resp.Streaming)

	// Send and receive a request synchronously.
	req = obs.NewGetStreamingStatusRequest()
	// Note that we create a new request,
	// because requests have IDs that must be unique.a
	// This will block until the response comes.
	resp, err = req.SendReceive(c)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("streaming:", resp.Streaming)

	// Respond to events by registering handlers.
	c.AddEventHandler("SwitchScenes", func(e obs.Event) {
		// Make sure to assert the actual event type.
		// The event is always a pointer.
		log.Println("new scene:", e.(*obs.SwitchScenesEvent).SceneName)
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
