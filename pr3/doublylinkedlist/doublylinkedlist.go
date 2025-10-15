package doublylinkedlist

type DoublyLinkedList[T any] struct {
	first *item[T]
	last  *item[T]
	size  int
}

type item[T any] struct {
	v    T
	next *item[T]
	prev *item[T]
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func (l *DoublyLinkedList[T]) Add(v T) {
	newItem := &item[T]{v: v}

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

func (l *DoublyLinkedList[T]) Get(idx int) T {
	if idx < 0 || idx >= l.size {
		panic("index out of bounds")
	}

	var current *item[T]
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

func (l *DoublyLinkedList[T]) Remove(idx int) {
	if idx < 0 || idx >= l.size {
		panic("index out of bounds")
	}

	var toRemove *item[T]
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

func (l *DoublyLinkedList[T]) Values() []T {
	result := make([]T, l.size)
	current := l.first
	for i := 0; i < l.size; i++ {
		result[i] = current.v
		current = current.next
	}
	return result
}