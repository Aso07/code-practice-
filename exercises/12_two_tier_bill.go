package main

import "fmt"

func TwoTierBill(name string, units int) string {
	bills := 0
	if units <= 50 {
		bills = units * 50
	}else if units > 50 {
		bills = (50 * 50) + ((units - 50) * 75 )
	}
	return fmt.Sprintf("%s owes %d", name, bills)
}
func main() {
	fmt.Println(TwoTierBill("Aso", 30))
	fmt.Println(TwoTierBill("Ido", 80))
}