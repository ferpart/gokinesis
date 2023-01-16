package starter

import (
	"fmt"
	"github.com/ferpart/gokinesis/domain"
)

func Start(k domain.IKinesis) {
	err := k.NewStream()
	if err != nil {
		panic(err)
	}

	fmt.Println("Kinesis stream started")
}
