package parallelism

import (
	"fmt"
	"sync"
)

// Неблокирующая функция, выполняющая разные операции со значением, пришедшем по разным каналам. Значение из stopChan прекратит работу.
func calculator1(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	out := make(chan int, 1)
	go func() {
		defer close(out)
		select {
		case value := <-firstChan:
			out <- value * value
		case value := <-secondChan:
			out <- 3 * value
		case <-stopChan:
		}
	}()
	return out
}

// Вызывается из main, управляет поступлением данных в каналы, запускает функцию calculator
func CalculatorFromMultiplChannelsRun() {
	firstChan := make(chan int)
	secondChan := make(chan int)
	defer func() {
		close(firstChan)
		close(secondChan)
	}()
	stopChan := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		out := calculator1(firstChan, secondChan, stopChan)
		fmt.Println(<-out)
	}()
	switch 2 {
	case 1:
		firstChan <- 2
	case 2:
		secondChan <- 2
	case 3:
		type emptyStruct struct{}
		myEmptyStruct := emptyStruct{}
		stopChan <- myEmptyStruct
	}
	wg.Wait()
}
