package main

import "fmt"

func Greet(name string) {
	fmt.Println("Hello " + name, " how are you")
}

func main() {
	Greet("Aso")
	Greet("Agene")
}