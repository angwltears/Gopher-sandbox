package main

import "fmt"

func task1() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [5]string{"a", "b", "c", "d", "e"}
	arr3 := [5]bool{true, false, true, false, true}
	fmt.Println(arr1[0])
	arr1[0] = 10
	fmt.Println(arr1, arr2, arr3)
}
