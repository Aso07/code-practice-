package main

import "fmt"

func TellSign(num int) string {
	if num > 0 {
		return "Positive"
	}
	if num < 0 {
		return "Negative"
	}
	return "Zero"
// }
// func main() {
// 	fmt.Println(TellSign(2))
// 	fmt.Println(TellSign(-1))
// 	fmt.Println(TellSign(0))
// }