package parallelism

import (
	"fmt"
	"sync"
	"time"
)

// Через канал arguments функция получит ряд чисел, а через канал done - сигнал о необходимости завершить работу.
// Когда сигнал о завершении работы будет получен, функция должна в выходной (возвращенный) канал отправить сумму полученных чисел.
func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	out := make(chan int, 1)
	go func() {
		defer close(out)
		sum := 0
		for {
			select {
			case value := <-arguments:
				sum += value
			case <-done:
				out <- sum
				return
			}
		}
	}()
	return out
}

// Вызывается из main, управляет поступлением данных в каналы, запускает функцию calculator
func CalculatorFromMultiplChannels2Run() {
	arguments := make(chan int)
	done := make(chan struct{})
	defer func() {
		close(arguments)
		close(done)
	}()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		out := calculator(arguments, done)
		fmt.Println(<-out)
	}()
	go func() {
		for range 15 {
			arguments <- 1
			time.Sleep(500 * time.Millisecond)
		}
	}()
	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Second)
		type emptyStruct struct{}
		myEmptyStruct := emptyStruct{}
		done <- myEmptyStruct
	}()
	wg.Wait()
}
