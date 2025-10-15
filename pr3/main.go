package main

import (
	"fmt"
	"pr3/binarytree"
	"pr3/doublylinkedlist"
	"pr3/matrixfiller"
	"pr3/romanconverter"
)

func main() {
	// Пример использования двусвязного списка
	fmt.Println("=== Doubly Linked List ===")
	list := doublylinkedlist.NewDoublyLinkedList[int]()
	list.Add(10)
	list.Add(20)
	list.Add(30)
	fmt.Println("List values:", list.Values())
	fmt.Println("Get index 1:", list.Get(1))
	list.Remove(1)
	fmt.Println("After remove index 1:", list.Values())

	// Пример использования бинарного дерева
	fmt.Println("\n=== Binary Tree ===")
	tree := binarytree.NewTree[int]()
	tree.Add(5)
	tree.Add(3)
	tree.Add(7)
	tree.Add(2)
	tree.Add(4)
	tree.Add(6)
	tree.Add(8)
	fmt.Println("Tree values (in-order):", tree.Values())
	tree.Remove(3)
	fmt.Println("After removing 3:", tree.Values())

	// Пример конвертации римских чисел
	fmt.Println("\n=== Roman to Integer ===")
	var input string
	fmt.Print("Введите римское число: ")
	fmt.Scanln(&input)
	result := romanconverter.RomanToInt(input)
	fmt.Printf("Арабское число: %d\n", result)

	// Пример заполнения двумерного массива
	fmt.Println("\n=== Unique 2D Matrix ===")
	matrix := matrixfiller.FillUnique2DArray(3, 4)
	fmt.Println("2D Array with unique numbers:")
	for _, row := range matrix {
		fmt.Println(row)
	}
}