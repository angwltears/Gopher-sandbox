package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(chan int)
	stop := make(chan bool)
	go func() {
		i := 0
		for {
			data <- i
			i++
			time.Sleep(500 * time.Millisecond)

		}
	}()
	go func() {
		time.Sleep(3 * time.Second)
		stop <- true
	}()
Loop:
	for {
		select {
		case <-stop:
			fmt.Println("Стоп, закрываю доступ")
			break Loop

		case v := <-data:
			fmt.Println("Работаю, ", v)
			time.Sleep(500 * time.Millisecond)

		}

	}
	fmt.Println("Программа завершилась корректно")
}
