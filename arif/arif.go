package main

import "fmt"

func task1() {
	fmt.Println("\n=== Задание 1 ===")

	var a, b float64
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	fmt.Printf("%.2f + %.2f = %.2f\n", a, b, a+b)
	fmt.Printf("%.2f - %.2f = %.2f\n", a, b, a-b)
	fmt.Printf("%.2f * %.2f = %.2f\n", a, b, a*b)
	if b != 0 {
		fmt.Printf("%.2f / %.2f = %.2f\n", a, b, a/b)
	} else {
		fmt.Println("Деление на ноль невозможно")
	}
}

func task2() {
	fmt.Println("\n=== Задание 2 ===")

	var year int
	fmt.Print("Введите год: ")
	fmt.Scanln(&year)

	isLeap := (year%400 == 0) || (year%100 != 0 && year%4 == 0)

	if isLeap {
		fmt.Printf("%d год — високосный\n", year)
	} else {
		fmt.Printf("%d год — не високосный\n", year)
	}
}

func main() {
	task1()
	task2()
}
