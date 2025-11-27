package main

import "fmt"

var phonebook = map[string]string{
	"Bob":   "8999",
	"Alice": "8800",
}

func main() {
	_, ok2 := phonebook["Bob"]
	if ok2 {
		delete(phonebook, "Bob")
	}

	val, ok := phonebook["Michael"]
	if !ok {
		fmt.Println("Не найден")
	} else {
		fmt.Println("Номер: ", val)
	}
	fmt.Println(phonebook)
}
