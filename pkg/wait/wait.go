package wait

import (
	"errors"
	"fmt"
	"time"
)

func For(f func() bool) error {
	return ForTime(f, 20)
}

func ForTime(f func() bool, seconds int) error {
	for i := 0; i < seconds; i++ {
		if f() == true {
			return nil
		}
		fmt.Println("Waiting...")
		time.Sleep(time.Second)
	}
	return errors.New("timeout")
}
