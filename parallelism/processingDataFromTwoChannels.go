package parallelism

import (
	"fmt"
	"sync"
	"time"
)

// n раз сделает следующее:
//   - прочитает по одному числу из каждого из двух каналов in1 и in2, назовем их x1 и x2.
//   - вычислит f(x1) + f(x2)
//   - запишет полученное значение в out
func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {

	mu1 := new(sync.Mutex)
	mu2 := new(sync.Mutex)
	res1m := make(map[int]int)
	res2m := make(map[int]int)

	// Ожидает выполнения fn, записывает результат в отображение
	fnWaiter := func(val, valNumber int, resMap map[int]int, mu *sync.Mutex) {
		res := fn(val)
		mu.Lock()
		resMap[valNumber] = res
		mu.Unlock()
	}
	// Получает данные из 1 канала и запускает горутины ожидания полезной работы
	go func() {
		for i := 0; i < n; i++ {
			val := <-in1
			go fnWaiter(val, i, res1m, mu1)
		}
	}()

	// Получает данные из 2 канала и запускает горутины ожидания полезной работы
	go func() {
		for i := 0; i < n; i++ {
			val := <-in2
			go fnWaiter(val, i, res2m, mu2)
		}
	}()

	// n раз ждёт появления результатов-слагаемых, формирует и отправляет ответ
	go func() {
		for i := range n {
			for {
				mu1.Lock()
				r1, ok1 := res1m[i]
				mu1.Unlock()
				mu2.Lock()
				r2, ok2 := res2m[i]
				mu2.Unlock()
				if ok1 && ok2 {
					out <- r1 + r2
					break
				}
			}
		}
	}()
}

var n int = 3
var wait chan string

// симулирует полезную работу для получения слагаемых, которая занимает различное время
func fn(x int) int {
	waitDur := <-wait
	duration, _ := time.ParseDuration(waitDur)
	fmt.Println("fnSleep ", duration, "for ", x)
	time.Sleep(duration)
	res := 10*x + x
	fmt.Println("fnRes ", res, waitDur)
	return res
}

// точка входа из main, тестирует функцию merge2Channels, отправляя и получая данные
func ProcessingDataFromTwoChannels() {
	wait = make(chan string, n*2)
	for _, duration := range []string{"800ms", "700ms", "600ms", "500ms", "400ms", "300ms"} {
		wait <- duration
	}
	in1 := make(chan int, 1)
	in2 := make(chan int, 1)
	out := make(chan int, 1)
	defer func() {
		close(in1)
		close(in2)
	}()
	var wg sync.WaitGroup

	merge2Channels(fn, in1, in2, out, n)

	wg.Add(3)
	// получает результаты от merge2Channels
	go func(out chan int) {
		defer wg.Done()
		for range n {
			value, ok := <-out
			fmt.Println("out ", value)
			if !ok {
				fmt.Println("out closed")
			}
		}
	}(out)

	go func(in1 chan int) {
		defer wg.Done()
		listIn1 := []int{1, 2, 3}
		for range n {
			in1 <- listIn1[0]
			fmt.Println("in1send ", listIn1[0])
			listIn1 = listIn1[1:]
		}
	}(in1)

	go func(in2 chan int) {
		defer wg.Done()
		listIn2 := []int{4, 5, 6}
		for range n {
			in2 <- listIn2[0]
			fmt.Println("in2send ", listIn2[0])
			listIn2 = listIn2[1:]
		}
	}(in2)

	wg.Wait()
}
