package parallelism

import (
	"fmt"
	"time"
)

func WFGwork() {
	time.Sleep(time.Second * 3)
	fmt.Println("Работа завершена")
}

func WaitinForGoroutina() {
	waitFunc := func() <-chan struct{} {
		done := make(chan struct{})
		go func() {
			WFGwork()
			close(done)
		}()
		return done
	}

	<-waitFunc()

}
