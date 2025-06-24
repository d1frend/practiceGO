package main

import "fmt"

type Transport interface {
	Move()
	Stop()
	GetName() string
}

type Car struct {
	Model string
	Speed int
}

func (c Car) Move() {
	fmt.Printf("%s движется со скоростью %d км/ч\n", c.Model, c.Speed)
}

func (c Car) Stop() {
	fmt.Printf("%s остановился\n", c.Model)
}

func (c Car) GetName() string {
	return c.Model
}

type Bicycle struct {
	Brand string
	Gear  int
}

func (b Bicycle) Move() {
	fmt.Printf("%s едет на %d передаче\n", b.Brand, b.Gear)
}

func (b Bicycle) Stop() {
	fmt.Printf("%s остановился\n", b.Brand)
}

func (b Bicycle) GetName() string {
	return b.Brand
}

func TestTransport(t Transport) {
	fmt.Println("\nТестируем транспорт:", t.GetName())
	t.Move()
	t.Stop()
}

func main() {
	car := Car{Model: "Toyota Camry", Speed: 60}
	bike := Bicycle{Brand: "Stels", Gear: 3}

	TestTransport(car)
	TestTransport(bike)

	garage := []Transport{car, bike}
	fmt.Println("\nВ гараже:")
	for _, vehicle := range garage {
		fmt.Println("-", vehicle.GetName())
	}
}
