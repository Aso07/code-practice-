package main

import(
	"fmt"
	"strings"
)
func StarTriangle(number int) {
	for i := 1; i <= number; i++ {
		fmt.Println(strings.Repeat("*", i))
	}
}
func main() {
	StarTriangle(5)
}