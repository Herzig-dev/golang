package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fillUnique2DArray(rows, cols int) [][]int {
	totalElements := rows * cols
	arr := make([][]int, rows)
	for i := range arr {
		arr[i] = make([]int, cols)
	}

	values := make([]int, totalElements)
	for i := 0; i < totalElements; i++ {
		values[i] = i + 1 
	}


	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(values), func(i, j int) {
		values[i], values[j] = values[j], values[i]
	})

	idx := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			arr[i][j] = values[idx]
			idx++
		}
	}

	return arr
}

func main() {
	rows := 3
	cols := 4

	arr := fillUnique2DArray(rows, cols)

	fmt.Println("Двумерный массив с уникальными числами:")
	for _, row := range arr {
		fmt.Println(row)
	}
}