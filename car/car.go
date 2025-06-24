package main

import "fmt"

type Engine struct {
	Type    string
	Power   int     //в л.с.
	Volume  float64 //в литрах
	IsTurbo bool
}

func (e Engine) PrintSpecs() {
	turboInfo := ""
	if e.IsTurbo {
		turboInfo = "с турбонаддувом"
	} else {
		turboInfo = "без турбонаддува"
	}
	fmt.Printf("Двигатель: %s, %.1f л, %d л.с. %s\n",
		e.Type, e.Volume, e.Power, turboInfo)
}

type Car struct {
	Make    string
	Model   string
	Year    int
	Engine  Engine
	Mileage int //пробег в км
}

func (c Car) PrintInfo() {
	fmt.Printf("Автомобиль: %s %s %d года\n", c.Make, c.Model, c.Year)
	fmt.Printf("Пробег: %d км\n", c.Mileage)
	fmt.Print("Характеристики двигателя: ")
	c.Engine.PrintSpecs()
}

func (c *Car) AddMileage(km int) {
	c.Mileage += km
	fmt.Printf("Пробег автомобиля %s %s увеличен на %d км. Теперь: %d км\n",
		c.Make, c.Model, km, c.Mileage)
}

func main() {
	engine1 := Engine{
		Type:    "Бензиновый",
		Power:   150,
		Volume:  2.0,
		IsTurbo: true,
	}

	car1 := Car{
		Make:    "Toyota",
		Model:   "Camry",
		Year:    2020,
		Engine:  engine1,
		Mileage: 35000,
	}

	car2 := Car{
		Make:    "BMW",
		Model:   "X5",
		Year:    2021,
		Engine:  Engine{"Дизельный", 249, 3.0, true},
		Mileage: 12000,
	}

	fmt.Println("=== Информация об автомобилях ===")
	car1.PrintInfo()
	fmt.Println()
	car2.PrintInfo()

	fmt.Println("\n=== Эксплуатация автомобилей ===")
	car1.AddMileage(500)
	car2.AddMileage(1000)

	fmt.Println("\n=== Тюнинг двигателя ===")
	car1.Engine.Power += 30
	fmt.Printf("Мощность двигателя %s %s увеличена до %d л.с.\n",
		car1.Make, car1.Model, car1.Engine.Power)
	car1.PrintInfo()
}
