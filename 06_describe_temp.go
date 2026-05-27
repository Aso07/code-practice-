package main

import "fmt"

func DescribeTemp(city string, temp int) string {
	if temp >= 35 {
		return city + " is hot"
	}
	if temp >= 25 {
		return city + " is warm"
	}
	if temp >= 15 {
		return city + " is mild"
	}
	return city + " is cold"
}
func main() {
	fmt.Println(DescribeTemp("Lagos", 38))
	fmt.Println(DescribeTemp("London", 12))
	fmt.Println(DescribeTemp("Abuja", 27))
	fmt.Println(DescribeTemp("Jos", 15))
}
