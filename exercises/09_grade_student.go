package main

import "fmt"

func GradeStudent(name string, score int) string {
	if score >= 70 {
		return fmt.Sprintf("%s you score A", name)
	}
	if score >= 60 {
		return fmt.Sprintf("%s you score B", name)
	}
	if score >= 50 {
		return fmt.Sprintf("%s you score C", name)
	}
	if score >= 40 {
		return fmt.Sprintf("%s you score D", name)
	}
	return fmt.Sprintf("%s you score F", name)
}
func main() {
	fmt.Println(GradeStudent("Aso",90))
	fmt.Println(GradeStudent("ido",65))
	fmt.Println(GradeStudent("ene",51))
	fmt.Println(GradeStudent("joy",49))
	fmt.Println(GradeStudent("bless",39))
}