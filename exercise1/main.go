package main

import "fmt"

type Mobile struct {
	Model string
	Brand string
}

func (m Mobile) Dial(phoneNumber string) string {
	return fmt.Sprintf("dial made to %s", phoneNumber)
}
func (m Mobile) SendMessage(phoneNumber string, text string) string {
	return fmt.Sprintf("text sent to %s. Message: [%s]", phoneNumber, text)
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

	nokia := Mobile{
		Brand: "Nokia",
		Model: "N73",
	}

	MakeDial(nokia, "0822xxxxx")
	MakeText(nokia, "0822xxxxx", "Hello World!")
}

func MakeDial(d Dialer, phoneNumber string) {
	fmt.Println(d.Dial(phoneNumber))
}

func MakeText(d Messenger, phoneNumber string, text string) {
	fmt.Println(d.SendMessage(phoneNumber, text))
}
