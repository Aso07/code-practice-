package main

import "fmt"

func BMICategory(name string, weight float64, height float64) string {
	BMI := weight / (height * height)
	if BMI < 18.5 {
		return fmt.Sprintf("%s is underweight", name)
	}
	if BMI < 25 {
		return fmt.Sprintf("%s is normal", name)
	}
	if BMI < 30 {
		return fmt.Sprintf("%s is overweight", name)
	}
	return fmt.Sprintf("%s is obese", name)
}
func main() {
	fmt.Println(BMICategory("Aso", 40.5, 1.2))
	fmt.Println(BMICategory("Aso", 70.5, 12.5))
	fmt.Println(BMICategory("Aso", 10.5, 5.5))
}
