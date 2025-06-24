package main

import (
	"fmt"
)

func task1() {
	fmt.Println("\n=== Задание 1: Калькулятор ===")

	var a, b float64
	var operator string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	fmt.Print("Введите оператор (+, -, *, /): ")
	fmt.Scanln(&operator)

	var result float64
	var err string

	if operator == "+" {
		result = a + b
	} else if operator == "-" {
		result = a - b
	} else if operator == "*" {
		result = a * b
	} else if operator == "/" {
		if b != 0 {
			result = a / b
		} else {
			err = "Ошибка: деление на ноль"
		}
	} else {
		err = "Ошибка: неизвестный оператор"
	}

	if err == "" {
		fmt.Printf("Результат: %.2f\n", result)
	} else {
		fmt.Println(err)
	}
}

func task2() {
	fmt.Println("\n=== Задание 2: День недели ===")

	var dayNumber int
	fmt.Print("Введите номер дня недели (1-7): ")
	fmt.Scanln(&dayNumber)

	var dayName string

	switch dayNumber {
	case 1:
		dayName = "Понедельник"
	case 2:
		dayName = "Вторник"
	case 3:
		dayName = "Среда"
	case 4:
		dayName = "Четверг"
	case 5:
		dayName = "Пятница"
	case 6:
		dayName = "Суббота"
	case 7:
		dayName = "Воскресенье"
	default:
		dayName = "Ошибка: введите число от 1 до 7"
	}

	fmt.Println(dayName)
}

func main() {
	task1()
	task2()
}
