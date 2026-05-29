package main

import(
	"fmt"
	"strings"
)
func HollowSquare(number int) {
	for i := 1; i <= number; i++ {
	if i == 1 {
		fmt.Println(strings.Repeat("*", number))
	}else if i == number {
		fmt.Println(strings.Repeat("*", number))
	}else{
	fmt.Println("*" + strings.Repeat(" ", number-2) + "*")
}
	}
}
func main() {
	HollowSquare(5)
}