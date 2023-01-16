package consumer

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ferpart/gokinesis/domain"
	"github.com/ferpart/gokinesis/pkg/consumer"
)

const timeout = time.Second * 5

func Consume(k domain.IKinesis) {
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
