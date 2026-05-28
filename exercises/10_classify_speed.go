package main

import "fmt"

func ClassifySpeed(name string, speed int) string {
	if speed > 120 {
		return fmt.Sprintf("%s Dangerously Fast", name)
	}
	if speed > 90 {
		return fmt.Sprintf("%s Speeding", name)
	}
	if speed > 60 {
		return fmt.Sprintf("%s Normal", name)
	}
	return fmt.Sprintf("%s Too Slow", name)
// }
// func main() {
// 	fmt.Println(ClassifySpeed("Aso", 150))
// 	fmt.Println(ClassifySpeed("ido", 100))
// 	fmt.Println(ClassifySpeed("kate", 80))
// 	fmt.Println(ClassifySpeed("joy", 50))
// }