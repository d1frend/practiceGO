package main

import (
	"fmt"
	"unsafe"
)

func basic_types() {
	var intExample int = -42
	var uintExample uint = 42
	var byteExample byte = 255
	var runeExample rune = 'Я'

	var float32Example float32 = 3.14
	var float64Example float64 = 3.1415926535

	var complex64Example complex64 = 1 + 2i
	var complex128Example complex128 = complex(3, 4)

	var boolExample bool = true

	var stringExample string = "Hello, Go!"

	fmt.Println("\n=== Задание 1 ===")
	fmt.Println("Целочисленные:")
	fmt.Printf("int: %d\n", intExample)
	fmt.Printf("uint: %d\n", uintExample)
	fmt.Printf("byte: %d\n", byteExample)
	fmt.Printf("rune: %U (символ: %c)\n", runeExample, runeExample)

	fmt.Println("\nЧисла с плавающей точкой:")
	fmt.Printf("float32: %.2f\n", float32Example)
	fmt.Printf("float64: %.5f\n", float64Example)

	fmt.Println("\nКомплексные числа:")
	fmt.Printf("complex64: %v\n", complex64Example)
	fmt.Printf("complex128: %v\n", complex128Example)

	fmt.Println("\nЛогический:")
	fmt.Printf("bool: %v\n", boolExample)

	fmt.Println("\nСтроки:")
	fmt.Printf("string: %s\n", stringExample)
}

func type_sizes() {
	fmt.Println("\n=== Задание 2 ===")
	fmt.Printf("int: %d байт\n", unsafe.Sizeof(int(0)))
	fmt.Printf("int8: %d байт\n", unsafe.Sizeof(int8(0)))
	fmt.Printf("int16: %d байт\n", unsafe.Sizeof(int16(0)))
	fmt.Printf("int32: %d байт\n", unsafe.Sizeof(int32(0)))
	fmt.Printf("int64: %d байт\n", unsafe.Sizeof(int64(0)))
	fmt.Printf("uint: %d байт\n", unsafe.Sizeof(uint(0)))
	fmt.Printf("uintptr: %d байт\n", unsafe.Sizeof(uintptr(0)))
	fmt.Printf("byte: %d байт\n", unsafe.Sizeof(byte(0)))
	fmt.Printf("rune: %d байт\n", unsafe.Sizeof(rune(0)))

	fmt.Printf("float32: %d байт\n", unsafe.Sizeof(float32(0)))
	fmt.Printf("float64: %d байт\n", unsafe.Sizeof(float64(0)))

	fmt.Printf("complex64: %d байт\n", unsafe.Sizeof(complex64(0)))
	fmt.Printf("complex128: %d байт\n", unsafe.Sizeof(complex128(0)))

	fmt.Printf("bool: %d байт\n", unsafe.Sizeof(bool(false)))

	fmt.Printf("string: %d байт (структура)\n", unsafe.Sizeof(""))
}

func main() {
	basic_types()
	type_sizes()
}
