package main

import "fmt"

type doublyLinkedList struct {
	first *item
	last  *item
	size  int
}

type item struct {
	v    any
	next *item
	prev *item
}

func newDoublyLinkedList() *doublyLinkedList {
	return &doublyLinkedList{}
}

// add - добавление значения в конец связного списка
func add(l *doublyLinkedList, v any) {
	newItem := &item{v: v}

	if l.size == 0 {
		l.first = newItem
		l.last = newItem
	} else {
		newItem.prev = l.last
		l.last.next = newItem
		l.last = newItem
	}

	l.size++
}

// get - получение значения по индексу из связанного списка
func get(l *doublyLinkedList, idx int) any {
	if idx < 0 || idx >= l.size {
		panic("index out of bounds")
	}

	var current *item
	if idx < l.size/2 {
		
		current = l.first
		for i := 0; i < idx; i++ {
			current = current.next
		}
	} else {
		
		current = l.last
		for i := l.size - 1; i > idx; i-- {
			current = current.prev
		}
	}

	return current.v
}

// remove - удаление значения по индексу из списка
func remove(l *doublyLinkedList, idx int) {
	if idx < 0 || idx >= l.size {
		panic("index out of bounds")
	}

	var toRemove *item
	if idx < l.size/2 {
		toRemove = l.first
		for i := 0; i < idx; i++ {
			toRemove = toRemove.next
		}
	} else {
		toRemove = l.last
		for i := l.size - 1; i > idx; i-- {
			toRemove = toRemove.prev
		}
	}

	if toRemove.prev != nil {
		toRemove.prev.next = toRemove.next
	} else {
		l.first = toRemove.next
	}

	if toRemove.next != nil {
		toRemove.next.prev = toRemove.prev
	} else {
		l.last = toRemove.prev
	}

	l.size--
}

// values - получение слайса значений из списка
func values(l *doublyLinkedList) []any {
	result := make([]any, l.size)
	current := l.first
	for i := 0; i < l.size; i++ {
		result[i] = current.v
		current = current.next
	}
	return result
}

func main() {
	list := newDoublyLinkedList()
	add(list, 17)
	add(list, 230)
	add(list, 564)

	fmt.Println("Values:", values(list)) 

	fmt.Println("Get index 1:", get(list, 1)) 

	remove(list, 1)
	fmt.Println("After remove index 1:", values(list)) 
}