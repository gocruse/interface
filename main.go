package main

import "fmt"

func main() {
	number := 2
	Cetak(number)
	kalimat := "Hello World!"
	Cetak(kalimat)
}

func Cetak(item interface{}) {
	fmt.Println("Item: ", item)
}
