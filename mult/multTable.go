package main

import (
	"fmt"
	"math"
)

func task1() {
	fmt.Println("\n=== Задание 1 ===")
	for i := 1; i <= 10; i++ {
		fmt.Printf("Таблица умножения на %d:\n", i)
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}
		fmt.Println()
	}
}

func task2() {
	fmt.Println("\n=== Задание 2 ===")
	fmt.Println("Простые числа до 100:")

	for num := 2; num <= 100; num++ {
		isPrime := true

		maxDivisor := int(math.Sqrt(float64(num))) + 1
		for i := 2; i < maxDivisor; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			fmt.Printf("%d ", num)
		}
	}
	fmt.Println()
}

func main() {
	task1()
	task2()
}
