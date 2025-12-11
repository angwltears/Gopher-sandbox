package main

import (
	"fmt"
	"time"
)

func main() {
	server1 := make(chan string)
	server2 := make(chan string)
	go func() {
		time.Sleep(1 * time.Millisecond)
		server1 <- "Ответ от сервера 1"
	}()
	go func() {
		time.Sleep(3 * time.Millisecond)
		server2 <- "Ответ от сервера 2"
	}()
	select {
	case msg := <-server1:
		fmt.Println(msg)

	case msg := <-server2:
		fmt.Println(msg)
	}
}
