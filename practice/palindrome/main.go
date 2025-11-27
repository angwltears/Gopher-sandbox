package main

import "fmt"

func main() {
	fmt.Println(palindrome("палитра"))

}
func palindrome(str string) bool {
	runes := []rune(str)
	pal := ""

	for i := len(runes) - 1; i >= 0; i-- {
		pal += string(runes[i])
	}
	if pal == str {
		return true
	} else {
		return false
	}
}
