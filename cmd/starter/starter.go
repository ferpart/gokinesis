package main

import (
	"fmt"

	"github.com/fpartidabc/gokinesis/pkg/kinesis"
)

const (
	hostname   = "http://localhost:4568"
	region     = "us-east-1"
	streamName = "qa-ssai-ad-tracking-regional"
)

func main() {
	k := kinesis.New(
		hostname,
		region,
		streamName,
	)

	err := k.NewStream()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Kinesis stream %s on host %s started", streamName, hostname)
}
