# go-obs-websocket

[![Build Status](https://travis-ci.com/christopher-dG/go-obs-websocket.svg?branch=master)](https://travis-ci.com/christopher-dG/go-obs-websocket)

`go-obs-websocket` provides client functionality for [`obs-websocket`](https://github.com/Palakis/obs-websocket).

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

	future, err := c.SendRequest(obs.NewGetStreamingStatusRequest())
	if err != nil {
		log.Fatal(err)
	}

	status := (<-future).(obs.GetStreamingStatusResponse)
	log.Println("streaming:", status.Streaming)

	c.AddEventHandler("Heartbeat", func(e obs.Event) {
		log.Println("profile:", e.(obs.HeartbeatEvent).CurrentProfile)
	})
	time.Sleep(time.Second * 10)
}
```
