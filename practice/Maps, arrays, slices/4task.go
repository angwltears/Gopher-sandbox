package main

import "fmt"

func Task4() {
	slots := make(map[string]float64)
	slots["A4"] = 500
	slots["B8"] = 920
	slots["D32"] = 1200
	slots["C5"] = 1500
	slots["E43"] = 300
	slots["F10"] = 400
	for k, v := range slots {
		if v < 500 {
			fmt.Println("Name: ", k)
		}
		if v < 900 {
			slots[k] = slots[k] * 0.9
		}
	}
	fmt.Println(slots)
}
