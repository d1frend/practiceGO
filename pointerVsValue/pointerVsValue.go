package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func byValue(val int) {
	val = 100
	fmt.Println("Inside byValue:", val)
}

func byPointer(ptr *int) {
	*ptr = 100
	fmt.Println("Inside byPointer:", *ptr)
}

func main() {
	original := 50

	fmt.Println("Original value before byValue:", original)
	byValue(original)
	fmt.Println("Original value after byValue:", original)

	fmt.Println("\nOriginal value before byPointer:", original)
	byPointer(&original)
	fmt.Println("Original value after byPointer:", original)

	p := Person{"Alice", 25}

	fmt.Println("\nOriginal person before byValue:", p)
	changePersonByValue(p)
	fmt.Println("Original person after byValue:", p)

	fmt.Println("\nOriginal person before byPointer:", p)
	changePersonByPointer(&p)
	fmt.Println("Original person after byPointer:", p)
}

func changePersonByValue(p Person) {
	p.Age = 30
	p.Name = "Bob"
	fmt.Println("Inside changePersonByValue:", p)
}

func changePersonByPointer(p *Person) {
	p.Age = 30
	p.Name = "Bob"
	fmt.Println("Inside changePersonByPointer:", *p)
}
