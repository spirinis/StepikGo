package parallelism

import (
	"fmt"
	"sync"
)

func SWGwork() {

}

func SyncWaitGroup() {
	wg := new(sync.WaitGroup)

	for range 10 {
		wg.Add(1)
		go func() {
			SWGwork()
			defer wg.Done()
			fmt.Println("Работа завершена")
		}()
	}
	wg.Wait()
	fmt.Println("Горутины завершили выполнение")
}
