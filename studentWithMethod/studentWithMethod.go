package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name      string
	BirthYear int
	Grades    []int
}

func (s Student) GetAge() int {
	currentYear := time.Now().Year()
	return currentYear - s.BirthYear
}

func (s Student) GetAverageGrade() float64 {
	if len(s.Grades) == 0 {
		return 0
	}
	sum := 0
	for _, grade := range s.Grades {
		sum += grade
	}
	return float64(sum) / float64(len(s.Grades))
}

func (s Student) GetStatus() string {
	avg := s.GetAverageGrade()
	switch {
	case avg >= 4.5:
		return "отличник"
	case avg >= 3.5:
		return "хорошист"
	case avg > 0:
		return "троечник"
	default:
		return "нет оценок"
	}
}

func (s Student) PrintInfo() {
	fmt.Printf("Студент: %s\n", s.Name)
	fmt.Printf("Возраст: %d лет\n", s.GetAge())
	fmt.Printf("Средний балл: %.1f\n", s.GetAverageGrade())
	fmt.Printf("Статус: %s\n", s.GetStatus())
}

func main() {
	student1 := Student{
		Name:      "Иван Иванов",
		BirthYear: 2000,
		Grades:    []int{5, 5, 4, 5},
	}

	student2 := Student{
		Name:      "Петр Петров",
		BirthYear: 2001,
		Grades:    []int{3, 4, 3, 4},
	}

	student1.PrintInfo()
	fmt.Println()
	student2.PrintInfo()
}
