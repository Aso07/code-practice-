package main

import "fmt"

func IsAdult(age int) bool {
	if age >= 18 {
		return true
	}
	return false
}
func main() {
	fmt.Println(IsAdult(20))
	fmt.Println(IsAdult(15))
	fmt.Println(IsAdult(18))
	fmt.Println(IsAdult(50))
}
