package main

import (
	"reflect"
	"sort"
	"testing"
)


func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func TestBubbleSort(t *testing.T) {
	input := []int{542, -565, 531, -294, -56, 14, 270, -51, -914, 605, -117, -768, 331, 708, -603, 84, -548, 579, 434, 751, 592, -349, 408, -602, 721, 909, 170, -432, -970, -171, -972, 316, 405, -676, -929, -795, -682, -646, 46, -609, -84, 180, -158, -662, -384, 854, -721, 39, 180, -197, -818, -946, -529, -555, -36, -853, -322, 540, -936, -919, 473, 978, 782, 586, 869, 333, -977, -548, -789, 988, -393, 807, -609, 997, 824, -480, -205, -576, 856, 494, 131, 40, -601, 467, 221, -640, 34, -220, 482, 948, 523, -27, -771, -914, 438, 957, 205, -411, -749, -723}

	inputCopy := make([]int, len(input))
	copy(inputCopy, input)

	bubbleSort(inputCopy)

	expected := make([]int, len(input))
	copy(expected, input)
	sort.Ints(expected)

	if !reflect.DeepEqual(inputCopy, expected) {
		t.Errorf("bubbleSort failed:\n got: %v\nwant: %v", inputCopy, expected)
	}

	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "reverse sorted",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "single element",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "duplicates",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5},
			expected: []int{1, 1, 2, 3, 4, 5, 5, 6, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)

			bubbleSort(inputCopy)
			if !reflect.DeepEqual(inputCopy, tt.expected) {
				t.Errorf("got %v, want %v", inputCopy, tt.expected)
			}
		})
	}
}