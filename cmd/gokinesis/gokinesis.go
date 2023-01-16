package main

import (
	"fmt"

	"github.com/spf13/pflag"

	"github.com/ferpart/gokinesis/internal/consumer"
	"github.com/ferpart/gokinesis/internal/starter"
	"github.com/ferpart/gokinesis/pkg/kinesis"
)

var operationType *string
var streamName *string
var hostname *string

func init() {
	operationType = pflag.StringP(
		"operation-type",
		"o",
		"start",
		"sets the operation to be run",
	)

	streamName = pflag.StringP(
		"stream-name",
		"s",
		"default",
		"sets the stream name to be created by the application. Defaults to \"default\"",
	)

	hostname = pflag.StringP(
		"hostname",
		"h",
		"http://localhost:4568",
		"sets the hostname in which the kinesis stream is located. Defaults to \"http://localhost:4568.\"",
	)
}

func main() {
	k := kinesis.New(
		*hostname,
		"us-east-1",
		*streamName,
	)

	switch *operationType {
	case "start":
		starter.Start(k)
	case "consume":
		consumer.Consume(k)
	default:
		panic(fmt.Errorf("invalid operation: %s", *operationType))
	}
}
