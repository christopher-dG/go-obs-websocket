package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"time"

	obs "github.com/christopher-dG/go-obs-websocket"
	"github.com/mitchellh/mapstructure"
	logging "github.com/op/go-logging"
)

var (
	host     = flag.String("host", "localhost", "obs-websocket hostname")
	port     = flag.Int("port", 4444, "obs-websocket port number")
	password = flag.String("password", "", "obs-websocket password")
	file     = flag.String("f", "", "JSON file to read requests from")
	quiet    = flag.Bool("q", false, "disable all output")
	logger   = logging.MustGetLogger("obsws")
)

type sleepRequest struct {
	Seconds     int    `yaml:"seconds"`
	RequestType string `yaml:"request-type"`
}

// sleepRequest needs to implement obs.Request.
func (r sleepRequest) ID() string   { return "" }
func (r sleepRequest) Type() string { return "" }

func main() {
	flag.Parse()
	if *file == "" {
		logger.Fatal("required argument 'file' not set")
	}

	if *quiet {
		logging.SetLevel(logging.ERROR, "obsws")
	} else {
		logging.SetLevel(logging.INFO, "obsws")
	}

	requests, err := readJSON()
	if err != nil {
		logger.Fatal("parsing input file failed:", err)
	}

	c := obs.Client{Host: *host, Port: *port, Password: *password}
	if err := c.Connect(); err != nil {
		logger.Fatal(err)
	}
	c.NoIDMode(true)
	defer func() {
		if err := c.Disconnect(); err != nil {
			logger.Warning(err)
		}
	}()

	sendRequests(c, requests)
}

func sendRequests(c obs.Client, requests []obs.Request) {
	for i, request := range requests {
		switch request.(type) {
		case sleepRequest:
			seconds := request.(sleepRequest).Seconds

			logger.Infof("request %d (Sleep): sleeping for %d second(s)", i, seconds)
			time.Sleep(time.Second * time.Duration(seconds))
		default:
			logger.Infof("request %d (%s): sending request", i, request.Type())
			future, err := c.SendRequest(request)
			if err != nil {
				logger.Warningf(
					"request %d (%s): sending request failed: %v",
					i, request.Type(), err,
				)
				continue
			}
			resp := <-future
			out, err := json.MarshalIndent(resp, "", "\t")
			if err != nil {
				logger.Warningf(
					"request %d (%s): couldn't marshal response: %v",
					i, request.Type(), err,
				)
				continue
			}
			logger.Infof("request %d (%s): response:\n%s", i, request.Type(), string(out))
		}
	}
}

func readJSON() ([]obs.Request, error) {
	bytes, err := ioutil.ReadFile(*file)
	if err != nil {
		return nil, err
	}
	raw := []map[string]interface{}{}
	if err = json.Unmarshal(bytes, &raw); err != nil {
		return nil, err
	}

	requests := []obs.Request{}
	for i, m := range raw {
		reqType := m["request-type"]
		if reqType == nil {
			logger.Warningf("request %d: invalid (no key 'request-type')", i)
			continue
		}

		if reqType == "Sleep" {
			request := &sleepRequest{}
			decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				TagName: "json",
				Result:  request,
			})
			if err != nil {
				logger.Warningf("request %d (Sleep): couldn't initialize decoder: %v", err)
				continue
			}
			if err = decoder.Decode(m); err != nil {
				logger.Warningf("request %d (Sleep): decoding failed: %v", i, err)
				continue
			}
			requests = append(requests, *request)
		} else {
			request := obs.ReqMap[reqType.(string)]
			if request == nil {
				logger.Warningf("request %d (%s): unknown request type", i, reqType)
				continue
			}

			decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				TagName:    "json",
				Result:     request,
				ZeroFields: true,
			})
			if err != nil {
				logger.Warningf(
					"request %d (%s): couldn't initialize decoder: %v",
					i, reqType, err,
				)
				continue
			}
			if err = decoder.Decode(m); err != nil {
				logger.Warningf("request %d (%s): decoding failed: %v", i, reqType, err)
				continue
			}
			requests = append(requests, request)
		}
	}

	if len(requests) == 0 {
		logger.Fatal("no requests were decoded")
	}

	return requests, nil
}
