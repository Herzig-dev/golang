package binarytree

import (
	"reflect"
	"testing"
)

func TestTree_Add_Values(t *testing.T) {
	tree := NewTree[int]()

	tree.Add(5)
	tree.Add(3)
	tree.Add(7)
	tree.Add(2)
	tree.Add(4)
	tree.Add(6)
	tree.Add(8)

	values := tree.Values()
	expected := []int{2, 3, 4, 5, 6, 7, 8}

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Values() = %v, want %v", values, expected)
	}
}

func TestTree_Remove(t *testing.T) {
	tree := NewTree[int]()

	tree.Add(5)
	tree.Add(3)
	tree.Add(7)
	tree.Add(2)
	tree.Add(4)
	tree.Add(6)
	tree.Add(8)

	tree.Remove(5)

	values := tree.Values()
	expected := []int{2, 3, 4, 6, 7, 8}

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Values() after Remove(5) = %v, want %v", values, expected)
	}

	tree.Remove(2)
	expected = []int{3, 4, 6, 7, 8}
	if !reflect.DeepEqual(tree.Values(), expected) {
		t.Errorf("Values() after Remove(2) = %v, want %v", tree.Values(), expected)
	}

	tree.Add(1)
	tree.Remove(3)
	expected = []int{1, 4, 6, 7, 8}
	if !reflect.DeepEqual(tree.Values(), expected) {
		t.Errorf("Values() after Remove(3) = %v, want %v", tree.Values(), expected)
	}
}

func TestTree_Empty(t *testing.T) {
	tree := NewTree[int]()

	if len(tree.Values()) != 0 {
		t.Errorf("New tree should be empty, got values: %v", tree.Values())
	}

	tree.Remove(10)
	if len(tree.Values()) != 0 {
		t.Errorf("Remove from empty tree should not add values, got: %v", tree.Values())
	}
}

func TestTree_Generic(t *testing.T) {
	tree := NewTree[string]()

	tree.Add("banana")
	tree.Add("apple")
	tree.Add("cherry")

	values := tree.Values()
	expected := []string{"apple", "banana", "cherry"}

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Generic Values() = %v, want %v", values, expected)
	}

	tree.Remove("banana")
	expected = []string{"apple", "cherry"}
	if !reflect.DeepEqual(tree.Values(), expected) {
		t.Errorf("Generic Values() after Remove = %v, want %v", tree.Values(), expected)
	}
}