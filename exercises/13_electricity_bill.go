package main

import "fmt"

func ElectricityBill(name string, units int) string {
	bill := 0
	if units <= 50 {
		bill = units * 50
	}else if units <= 150 {
		bill = (50 * 50) + ((units - 50)* 75)
	}else if units > 150 {
		bill = (50 * 50) + (100 * 75) + (units - 150)* 100
	}
	return fmt.Sprintf("%s owes %d", name, bill)
}
func main() {
	fmt.Println(ElectricityBill("Aso", 30))
	fmt.Println(ElectricityBill("Ido", 80))
	fmt.Println(ElectricityBill("joy", 200))
}