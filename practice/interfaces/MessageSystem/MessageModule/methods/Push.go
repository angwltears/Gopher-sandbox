package methods

import "fmt"

type Push struct {
	Message  string
	Receiver string
}

func (e Push) Send() {
	fmt.Println("Push message:", e.Message)
	fmt.Println("Push receiver: ", e.Receiver)
}
