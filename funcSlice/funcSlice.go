package main

import (
	"fmt"
	"sort"
)

func FindElement(slice []int, target int) (int, bool) {
	for i, v := range slice {
		if v == target {
			return i, true
		}
	}
	return -1, false
}

func SortSlice(slice []int) []int {
	sorted := make([]int, len(slice))
	copy(sorted, slice)
	sort.Ints(sorted)
	return sorted
}

func FilterEvenNumbers(slice []int) []int {
	var result []int
	for _, v := range slice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

func PrintSlice(slice []int) {
	fmt.Printf("%v\n", slice)
}

func main() {
	numbers := []int{5, 3, 8, 1, 2, 7, 4, 6}

	fmt.Println("Исходный срез:")
	PrintSlice(numbers)

	target := 8
	if index, found := FindElement(numbers, target); found {
		fmt.Printf("Элемент %d найден на позиции %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден\n", target)
	}

	sorted := SortSlice(numbers)
	fmt.Println("Отсортированный срез:")
	PrintSlice(sorted)

	evens := FilterEvenNumbers(numbers)
	fmt.Println("Только четные числа:")
	PrintSlice(evens)
}
