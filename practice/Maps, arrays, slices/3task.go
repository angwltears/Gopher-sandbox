package main

import "fmt"

type Dog struct {
	name   string
	rate   int
	gentle bool
}

func task4() {
	dog1 := Dog{"Kasper", 10, true}
	dog2 := Dog{"Alex", 5, false}
	dog3 := Dog{"Jeremy", 2, false}
	list := []Dog{dog1, dog2, dog3}
	fmt.Println("before: ", list)
	for i := 0; i < len(list); i++ {
		if list[i].gentle {
			list[i].rate += 1
		}
	}
	fmt.Println("after: ", list)
}
