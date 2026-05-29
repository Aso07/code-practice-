package main

import "fmt"

func ShippingCost(name string, weight int) string {
	cost := 0
	if weight <= 5 {
		cost = weight * 500
	}else if weight <= 15 {
		cost = (5 * 500) + ((weight - 5 )* 350)
	}else {
		cost = (5 * 500) + (10 * 350) + (weight - 15)* 200
	}
	return fmt.Sprintf("%s cost is %d", name, cost)
}
func main() {
	fmt.Println(ShippingCost("Aso",3))
	fmt.Println(ShippingCost("Ido",10))
	fmt.Println(ShippingCost("Joy",20))
}