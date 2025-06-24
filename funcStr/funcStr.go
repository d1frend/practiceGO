package main

import "fmt"

func FindLongestString(strings ...string) string {
	if len(strings) == 0 {
		return ""
	}

	longest := strings[0]
	for _, s := range strings[1:] {
		if len(s) > len(longest) {
			longest = s
		}
	}
	return longest
}

func main() {
	fmt.Println("Самая длинная строка:", FindLongestString("один", "два", "три", "четыре"))
	fmt.Println("Самая длинная строка:", FindLongestString("яблоко", "банан", "апельсин"))
	fmt.Println("Самая длинная строка:", FindLongestString("короткая", "очень длинная строка", "средняя"))

	fmt.Println("Самая длинная строка:", FindLongestString())

	words := []string{"Go", "Python", "JavaScript", "Java"}
	fmt.Println("Самая длинная строка:", FindLongestString(words...))
}
