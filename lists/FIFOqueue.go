package lists

import (
	"container/list"
	"fmt"
)

// добавление элемента
func Push(elem interface{}, queue *list.List) {
	queue.PushBack(elem)
}

// Из Pop() удаляется первое записанное число.
func Pop(queue *list.List) interface{} {
	return queue.Remove(queue.Front())
}

// печать очереди в одну строку без пробелов
func printQueue(queue *list.List) {
	elem := queue.Front()
	for range queue.Len() {
		fmt.Print(elem.Value)
		elem = elem.Next()
	}
}

func FIFOqueue() {
	myList := list.New()
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	myList.PushBack(4)
	myList.PushBack(5)

	printQueue(myList)
	fmt.Println()
	fmt.Println(Pop(myList))
	printQueue(myList)
	fmt.Println()
	Push(1, myList)
	printQueue(myList)
	fmt.Println()
}
