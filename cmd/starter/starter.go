package main

import (
	"errors"
	"fmt"

	"github.com/spf13/pflag"

	"github.com/ferpart/gokinesis/pkg/kinesis"
)

var streamName *string
var hostname *string

func init() {
	streamName = pflag.StringP(
		"stream-name",
		"s",
		"",
		"Required: sets the stream name to be created by the application",
	)

	hostname = pflag.StringP(
		"hostname",
		"h",
		"http://localhost:4568",
		"Sets the hostname in which the kinesis stream is located. Defaults to \"http://localhost:4568.\"",
	)
}

func main() {
	if *streamName == "" {
		panic(errors.New("no stream name provided"))
	}

	k := kinesis.New(
		*hostname,
		"us-east-1",
		*streamName,
	)

	err := k.NewStream()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Kinesis stream %s on host %s started", *streamName, *hostname)
}
