package main

import "fmt"

type LivingThing interface {
	Run(length int) string
}

// Define interface
type Human interface {
	LivingThing
	Cry() string
}

type Student struct {
	Name string
}

func (s Student) Run(length int) string {
	return fmt.Sprintf("%s run for %d meters", s.Name, length)
}

func (s Student) Cry() string {
	return fmt.Sprintf("%s cry like a crazy", s.Name)
}

type Dog struct {
	Name string
}

func (d Dog) Run(length int) string {
	return fmt.Sprintf("%s the Dog, run for %d meters", d.Name, length)
}

func main() {

	studentJohn := Student{
		Name: "John",
	}

	MakeCry(studentJohn)
	MakeHumanRun(studentJohn, 20)

	hachiko := Dog{
		Name: "Hachiko",
	}

	MakeRun(hachiko, 10)
	MakeCry(hachiko) // This will throw error on compile
}

func MakeRun(o LivingThing, length int) {
	fmt.Println(o.Run(length))
}

func MakeHumanRun(human Human, length int) {
	fmt.Println(human.Run(length))
}

func MakeCry(human Human) {
	fmt.Println(human.Cry())
}
