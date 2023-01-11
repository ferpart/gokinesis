package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/fpartidabc/gokinesis/internal/common2csv"
	"github.com/fpartidabc/gokinesis/internal/consume"
	"github.com/fpartidabc/gokinesis/pkg/kinesis"
)

const (
	hostname   = "http://localhost:4568"
	region     = "us-east-1"
	streamName = "qa-ssai-ad-tracking-regional"

	//timeout = time.Second
)

func main() {
	k := kinesis.New(
		hostname,
		region,
		streamName,
	)

	stop := make(chan os.Signal, 1)
	signal.Notify(
		stop,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)

	stopConsume := make(chan bool, 1)
	storeChan := make(chan *consume.CommonMap, 1)

	go consume.Consume(k, stopConsume, storeChan)

	<-stop

	stopConsume <- true

	if err := common2csv.Common2CSV(<-storeChan); err != nil {
		panic(err)
	}
}
