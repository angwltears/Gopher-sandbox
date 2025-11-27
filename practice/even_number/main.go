package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("before: ", slice)
	fmt.Println("after: ", even(slice))
}
func even(slc []int) []int {
	newslc := make([]int, 0)
	for _, v := range slc {
		if v%2 == 0 {
			newslc = append(newslc, v)
		}
	}
	return newslc
}
