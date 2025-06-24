package main

import "fmt"

func main() {
	grades := make(map[string]int)

	grades["Максон"] = 5
	grades["Тимоха"] = 5
	grades["Матвей"] = 4
	grades["Андрюха"] = 3
	grades["Вовчик"] = 3

	fmt.Println("Оценка Вовчика:", findGrade(grades, "Вовчик"))
	fmt.Println("Оценка Петра:", findGrade(grades, "Петр")) //такого нет

	removeGrade(grades, "Андрюха")
	fmt.Println("Оценка Андрюхи после удаления:", findGrade(grades, "Андрюха"))

	fmt.Println("Все оценки:", grades)
}

func addGrade(grades map[string]int, name string, grade int) {
	grades[name] = grade
}

func findGrade(grades map[string]int, name string) int {
	grade, exists := grades[name]
	if !exists {
		return -1
	}
	return grade
}

func removeGrade(grades map[string]int, name string) {
	delete(grades, name)
}
