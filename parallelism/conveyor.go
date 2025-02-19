package parallelism

import "fmt"

func removeDuplicates(inputStream, outputStream chan string) {
	defer close(outputStream)
	var prevStr string
	firstIteration := true
	for {
		currentStr, ok := <-inputStream
		if !ok {
			break
		}
		if currentStr != prevStr || firstIteration {
			if firstIteration {
				firstIteration = false
			}
			outputStream <- currentStr
			prevStr = currentStr
		}
	}
}

func send(inputStream chan string) {
	defer func() {
		close(inputStream)
		fmt.Println("inputStream закрыт")
	}()
	for _, element := range []string{"ab", "aaaaaaaa", "b", "cc", "ccc", "d", "d", "e"} {
		inputStream <- element
		fmt.Println("Отправлено:", element)
	}
}

func Conveyor() {
	inputStream := make(chan string)
	outputStream := make(chan string)

	go removeDuplicates(inputStream, outputStream)
	go send(inputStream)

	for result := range outputStream {
		fmt.Println("Получено:", result)
	}
}
