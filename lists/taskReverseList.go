package lists

import (
	"container/list"
	"fmt"
)

// ReverseList - функция для реверса списка
func ReverseList(l *list.List) *list.List {
	reversedList := list.New()
	for elem := l.Front(); elem != nil; elem = elem.Next() {
		reversedList.PushFront(elem.Value)
	}
	return reversedList
}

// печать очереди в одну строку без пробелов
func print(queue *list.List) {
	elem := queue.Front()
	for range queue.Len() {
		fmt.Print(elem.Value, " ")
		elem = elem.Next()
	}
	fmt.Println()
}

func TaskReverseList() {
	myList := list.New()
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	myList.PushBack(4)
	myList.PushBack(5)
	print(myList)
	print(ReverseList(myList))
}
