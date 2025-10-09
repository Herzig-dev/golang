package main

import (
	"errors"
	"fmt"
)

type stack struct {
	s    []any // слайс в котором хранятся значения в стеке
	head int  // индекс головы стека
}

func newStack(size int) *stack {
	return &stack{
		s:    make([]any, size),
		head: -1,
	}
}

// push - добавление в стек значения
func push(s *stack, v any) error {
	if s.head+1 >= len(s.s) {
		return errors.New("stack overflow")
	}
	s.head++
	s.s[s.head] = v
	return nil
}

// pop - получения значения из стека и его удаление из вершины
func pop(s *stack) (any, error) {
	if s.head == -1 {
		return nil, errors.New("stack underflow")
	}
	value := s.s[s.head]
	s.head--
	return value, nil
}

// peek - просмотр значения на вершине стека
func peek(s *stack) (any, error) {
	if s.head == -1 {
		return nil, errors.New("stack is empty")
	}
	return s.s[s.head], nil
}

func main() {
	stack := newStack(5)

	push(stack, 15)
	push(stack, 76)
	push(stack, 245)

	top, err := peek(stack)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Top element:", top) 
	}

	value, _ := pop(stack)
	fmt.Println("Popped:", value) 

	top, err = peek(stack)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Top after pop:", top) 
	}
}