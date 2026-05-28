package main

import "fmt"

func FlatBill(name string, units int) string {
	bills := units * 50 
		return fmt.Sprintf("%s owes %d", name, bills)
	
// }
// func main() {
// 	fmt.Println(FlatBill("Aso", 30))
// 	fmt.Println(FlatBill("Ido", 80))
// 	fmt.Println(FlatBill("Joy", 200))
// }