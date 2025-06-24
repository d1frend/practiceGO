package main

import (
	"fmt"
	"time"
)

func main() {
	name := "Danek"
	currentTime := time.Now()

	fmt.Println("Hello, World!")
	fmt.Printf("Имя: %s\n", name)
	fmt.Printf("Текущая дата: %s\n", currentTime.Format("02.01.2006"))
}
