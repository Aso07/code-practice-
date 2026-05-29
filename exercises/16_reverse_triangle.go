package main

import(
	"fmt"
	"strings"
)
func ReverseTriangle(number int) {
	for i := number; i >= 1; i-- {
		fmt.Println(strings.Repeat("*", i))
	}
}
func main() {
	ReverseTriangle(5)
}