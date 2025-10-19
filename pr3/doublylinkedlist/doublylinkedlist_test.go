package doublylinkedlist

import (
	"reflect"
	"testing"
)

func TestDoublyLinkedList_Add_Get_Values(t *testing.T) {
	list := NewDoublyLinkedList[int]()

	list.Add(10)
	list.Add(20)
	list.Add(30)

	values := list.Values()
	expected := []int{10, 20, 30}

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Values() = %v, want %v", values, expected)
	}

	if val := list.Get(1); val != 20 {
		t.Errorf("Get(1) = %v, want %v", val, 20)
	}

	if val := list.Get(0); val != 10 {
		t.Errorf("Get(0) = %v, want %v", val, 10)
	}

	if val := list.Get(2); val != 30 {
		t.Errorf("Get(2) = %v, want %v", val, 30)
	}
}

func TestDoublyLinkedList_Remove(t *testing.T) {
	list := NewDoublyLinkedList[int]()

	list.Add(10)
	list.Add(20)
	list.Add(30)
	list.Add(40)

	list.Remove(1)

	values := list.Values()
	expected := []int{10, 30, 40}

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Values() after Remove(1) = %v, want %v", values, expected)
	}

	list.Remove(0)
	expected = []int{30, 40}
	if !reflect.DeepEqual(list.Values(), expected) {
		t.Errorf("Values() after Remove(0) = %v, want %v", list.Values(), expected)
	}

	list.Remove(1)
	expected = []int{30}
	if !reflect.DeepEqual(list.Values(), expected) {
		t.Errorf("Values() after Remove(last) = %v, want %v", list.Values(), expected)
	}

	list.Remove(0)
	expected = []int{}
	if !reflect.DeepEqual(list.Values(), expected) {
		t.Errorf("Values() after removing last element = %v, want %v", list.Values(), expected)
	}
}

func TestDoublyLinkedList_Empty(t *testing.T) {
	list := NewDoublyLinkedList[int]()

	if len(list.Values()) != 0 {
		t.Errorf("New list should be empty, got values: %v", list.Values())
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Get on empty list should panic")
		}
	}()
	list.Get(0) 
}

func TestDoublyLinkedList_Generic(t *testing.T) {
	list := NewDoublyLinkedList[string]()

	list.Add("hello")
	list.Add("world")

	values := list.Values()
	expected := []string{"hello", "world"}

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Generic Values() = %v, want %v", values, expected)
	}

	if val := list.Get(1); val != "world" {
		t.Errorf("Generic Get(1) = %v, want %v", val, "world")
	}
}