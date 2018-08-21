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

	"github.com/christopher-dG/go-obs-websocket"
)

func main() {
	// Connect a client.
	c := obsws.Client{Host: "localhost", Port: 4444}
	if err := c.Connect(); err != nil {
		log.Fatal(err)
	}
	defer c.Disconnect()

	// Send and receive a request asynchronously.
	req := obsws.NewGetStreamingStatusRequest()
	if err := req.Send(c); err != nil {
		log.Fatal(err)
	}
	// This will block until the response comes (potentially forever).
	resp, err := req.Receive()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("streaming:", resp.Streaming)

	// Set the amount of time we can wait for a response.
	obsws.SetReceiveTimeout(time.Second)

	// Send and receive a request synchronously.
	req = obsws.NewGetStreamingStatusRequest()
	// Note that we create a new request,
	// because requests have IDs that must be unique.a
	// This will block for up to two seconds, since we set a timeout.
	resp, err = req.SendReceive(c)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("streaming:", resp.Streaming)

	// Respond to events by registering handlers.
	c.AddEventHandler("SwitchScenes", func(e obsws.Event) {
		// Make sure to assert the actual event type.
		log.Println("new scene:", e.(obs.SwitchScenesEvent).SceneName)
	})

	time.Sleep(time.Second * 10)
}
```
