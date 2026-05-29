package main

import(
	"fmt"
	"strings"
)
func PrintSquare(number int) {
	for i := 1; i <= number; i++ {
		fmt.Println(strings.Repeat("*", number))
	}
}
func main() {
	PrintSquare(4)
}