package main

import "fmt"

type Mobile struct {
	Model string
	Brand string
}

type Dialer interface {
	Dial(phoneNumber string) string
}

type Messenger interface {
	SendMessage(phoneNumber string, text string) string
}

func main() {
	// Task:
	// - Create a Mobile item from the struct
	// - Make a Dial -> Call the MakeDial function
	// - Send a text message -> call the MakeText function

}

func MakeDial(d Dialer, phoneNumber string) {
	fmt.Println(d.Dial(phoneNumber))
}

func MakeText(d Messenger, phoneNumber string, text string) {
	fmt.Println(d.SendMessage(phoneNumber, text))
}
