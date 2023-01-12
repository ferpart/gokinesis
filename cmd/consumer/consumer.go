package main

import (
	"errors"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/pflag"

	"github.com/ferpart/gokinesis/pkg/consumer"
	"github.com/ferpart/gokinesis/pkg/kinesis"
)

const timeout = time.Second * 5

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

	stop := make(chan os.Signal, 1)
	signal.Notify(
		stop,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)

	stopConsume := make(chan bool, 1)

	c := consumer.New(k)
	go func() {
		if err := c.Consume(stopConsume); err != nil {
			panic(err)
		}
	}()

	<-stop
	stopConsume <- true

	time.Sleep(timeout)
}
