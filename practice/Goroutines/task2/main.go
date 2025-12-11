package main

import (
	"fmt"
	"strconv"
)

func main() {
	messages := make(chan string)

	go Send(messages, 1)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println("----------------------------------------------------")
	messagesbuffer := make(chan string, 2)
	fmt.Println(len(messagesbuffer))
	messagesbuffer <- "Сообщение 1"
	messagesbuffer <- "Сообщение 2"
	fmt.Println(<-messagesbuffer)
	fmt.Println(<-messagesbuffer)
}
func Send(messages chan string, n int) {
	messages <- "Важное сообщение:" + strconv.Itoa(n)
	n++
	messages <- "Важное сообщение:" + strconv.Itoa(n)

}
