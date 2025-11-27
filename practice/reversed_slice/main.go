package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("before: ", slice)
	fmt.Println("after: ", reverse(slice))
}
func reverse(slc []int) []int {
	newslc := make([]int, 0)
	for i := len(slc) - 1; i >= 0; i-- {
		newslc = append(newslc, slc[i])
	}
	return newslc
}
