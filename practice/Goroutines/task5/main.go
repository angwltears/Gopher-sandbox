package main

import (
	"fmt"
	"time"
)

func main() {
	database := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		database <- "Результат найден"
	}()
	select {

	case msg := <-database:
		fmt.Println(msg)

	case <-time.After(2 * time.Second):
		fmt.Println("Ошибка, Таймаут")
	}
}
