package main

import "fmt"

type tree struct {
	head *node
}

type node struct {
	left, right *node
	v           int
}

func newTree() *tree {
	return &tree{}
}

// add - добавление значения в дерево
func add(t *tree, v int) {
	t.head = insert(t.head, v)
}

func insert(n *node, v int) *node {
	if n == nil {
		return &node{v: v}
	}

	if v < n.v {
		n.left = insert(n.left, v)
	} else if v > n.v {
		n.right = insert(n.right, v)
	}

	return n
}

// remove - удаление значения из дерева
func remove(t *tree, v int) {
	t.head = deleteNode(t.head, v)
}

func deleteNode(n *node, v int) *node {
	if n == nil {
		return nil
	}

	if v < n.v {
		n.left = deleteNode(n.left, v)
	} else if v > n.v {
		n.right = deleteNode(n.right, v)
	} else {
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}

		minNode := findMin(n.right)
		n.v = minNode.v
		n.right = deleteNode(n.right, minNode.v)
	}

	return n
}

func findMin(n *node) *node {
	for n.left != nil {
		n = n.left
	}
	return n
}

// values - получение отсортированного слайса значений из дерева
func values(t *tree) []int {
	result := []int{}
	inOrder(t.head, &result)
	return result
}

func inOrder(n *node, result *[]int) {
	if n != nil {
		inOrder(n.left, result)
		*result = append(*result, n.v)
		inOrder(n.right, result)
	}
}

func main() {
	t := newTree()
	add(t, 5)
	add(t, 3)
	add(t, 7)
	add(t, 2)
	add(t, 4)
	add(t, 6)
	add(t, 8)

	fmt.Println("Values:", values(t)) 

	remove(t, 3)
	fmt.Println("After removing 3:", values(t)) 
}