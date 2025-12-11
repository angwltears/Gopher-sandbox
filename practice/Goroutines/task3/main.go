package main

import "fmt"

func main() {
	nums := make(chan int)
	go generator(nums)
	printer(nums)
}
func generator(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}
func printer(ch <-chan int) {
	for n := range ch {
		fmt.Println(n)
	}
}
