package main

import (
	"flag"

	obs "github.com/christopher-dG/go-obs-websocket"
	logging "github.com/op/go-logging"
)

var (
	host     = flag.String("host", "localhost", "obs-websocket hostname")
	port     = flag.Int("port", 4444, "obs-websocket port number")
	password = flag.String("password", "", "obs-websocket password")
	verbose  = flag.Bool("v", false, "enable debug output")
	logger   = logging.MustGetLogger("obsws")
)

func main() {
	flag.Parse()
	if *verbose {
		logging.SetLevel(logging.DEBUG, "obsws")
	} else {
		logging.SetLevel(logging.INFO, "obsws")
	}

	c := obs.Client{Host: *host, Port: *port, Password: *password}
	if err := c.Connect(); err != nil {
		logger.Fatal(err)
	}
	defer func() {
		if err := c.Disconnect(); err != nil {
			logger.Warning(err)
		} else {
			logger.Debug("logged out")
		}
	}()
}
