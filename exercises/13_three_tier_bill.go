package main

import "fmt"

func ThreeTierBill(name string, units int) string {
	bill := 0
	if units <= 50 {
		bill = units * 50
	}else if units <= 150 {
		bill = (50 * 50) + ((units * 50)* 75)
	}else if units > 150 {
		bill = (50 * 50) + ((units * 50)* 75)* 100
	}
	return fmt.Sprintf("%s owes %d", name, bill)
}
func main() {
	fmt.Println(ThreeTierBill("Aso", 30))
	fmt.Println(ThreeTierBill("Ido", 80))
	fmt.Println(ThreeTierBill("joy", 200))
}