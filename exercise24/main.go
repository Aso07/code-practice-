package main

import(
	"fmt"
)
func MiddleLetter(s string) string {
	return s [len(s) / 2 : len(s)/2+1]
}
func main() {
	fmt.Println(MiddleLetter("CAT"))
	fmt.Println(MiddleLetter("HELLO"))
	fmt.Println(MiddleLetter("GOLANG"))
}