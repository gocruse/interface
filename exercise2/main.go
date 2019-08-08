package main

import "fmt"

type Car struct {
	Model     string
	Brand     string
	TotalDoor uint
}

func (c Car) Move(length int) string {
	return fmt.Sprintf("The %s %s car has moved for %d meters", c.Brand, c.Model, length)
}

type Bicycle struct {
	Model string
	Brand string
}

func (b Bicycle) Move(length int) string {
	return fmt.Sprintf("The %s %s bicycle rided and moved for %d meters", b.Brand, b.Model, length)
}

type Mover interface {
	Move(lenth int) string
}

func main() {
	car := Car{
		Brand:     "Honda",
		Model:     "Civic",
		TotalDoor: 4,
	}

	bicycle := Bicycle{
		Brand: "BMX",
		Model: "X6",
	}

	MakeMove(car, 2)
	MakeMove(bicycle, 2)

	// Task:
	//  The whole code is already worked well. But it's look like there a few redundant function.
	// - Make an abstraction from this two struct
	// - See the pattern, and create abstraction
	// - Feel free to create a new function. Stand for the DRY (Don't Repeat Yourself) principle
}

func MakeMove(m Mover, length int) {
	fmt.Println(m.Move(length))
}
