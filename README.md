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
