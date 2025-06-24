package main

import "fmt"

type Student struct {
	Name     string
	Age      int
	Course   int
	AvgGrade float64
}

func NewStudent(name string, age, course int, avgGrade float64) Student {
	return Student{
		Name:     name,
		Age:      age,
		Course:   course,
		AvgGrade: avgGrade,
	}
}

func (s Student) PrintInfo() {
	fmt.Printf("Студент: %s\n", s.Name)
	fmt.Printf("Возраст: %d\n", s.Age)
	fmt.Printf("Курс: %d\n", s.Course)
	fmt.Printf("Средний балл: %.2f\n", s.AvgGrade)
}

func (s *Student) Promote() {
	s.Course++
	fmt.Printf("%s переведен на %d курс\n", s.Name, s.Course)
}

func (s *Student) UpdateGrade(newGrade float64) {
	s.AvgGrade = newGrade
	fmt.Printf("Новый средний балл для %s: %.2f\n", s.Name, s.AvgGrade)
}

func main() {
	student1 := NewStudent("Иван Петров", 20, 2, 4.5)
	var student2 Student
	student2.Name = "Мария Сидорова"
	student2.Age = 21
	student2.Course = 3
	student2.AvgGrade = 4.8

	fmt.Println("=== Информация о студентах ===")
	student1.PrintInfo()
	fmt.Println()
	student2.PrintInfo()

	fmt.Println("\n=== Обновление данных ===")
	student1.Promote()
	student2.UpdateGrade(4.9)

	students := []Student{student1, student2}
	fmt.Println("\n=== Список всех студентов ===")
	for _, s := range students {
		s.PrintInfo()
		fmt.Println()
	}
}
