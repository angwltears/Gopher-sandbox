package main

import "fmt"

func main() {
	fmt.Println("Начало")
	done := make(chan bool)
	go func() {
		fmt.Println("Привет из горутины!")
		done <- true
	}()
	ok := <-done
	fmt.Println("Возвращенное значение: ", ok)
	fmt.Println("Конец.")
}
