package main

import (
	"fmt"
	"strings"
)

func task1() {
	fmt.Println("\n=== Задание 1 ===")

	str := "GoLang — мощный язык программирования! Go Go GO"

	fmt.Println("Верхний регистр:", strings.ToUpper(str))
	fmt.Println("Нижний регистр:", strings.ToLower(str))

	fmt.Println("Начинается с 'Go'?", strings.HasPrefix(str, "Go"))
	fmt.Println("Заканчивается на '!'?", strings.HasSuffix(str, "!"))

	substr := "Go"
	fmt.Println("Содержит 'язык'?", strings.Contains(str, "язык"))
	fmt.Printf("'%s' встречается %d раз(а)\n", substr, strings.Count(str, substr))

	newStr := strings.ReplaceAll(str, "Go", "GOLANG")
	fmt.Println("После замены:", newStr)
}

func task2() {
	fmt.Println("\n=== Задание 2 ===")

	sentence := "Раз,два,три,четыре,пять"
	fmt.Println(sentence)

	words := strings.Split(sentence, ",")
	for i, word := range words {
		fmt.Printf("Слово %d: %s\n", i+1, word)
	}

	joined := strings.Join(words, " | ")
	fmt.Println("\nСоединённая строка:", joined)
}

func main() {
	task1()
	task2()
}
