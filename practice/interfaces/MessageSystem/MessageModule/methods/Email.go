package methods

import "fmt"

type Email struct {
	Message  string
	Receiver string
}

func (e Email) Send() {
	fmt.Println("Email message:", e.Message)
	fmt.Println("Email receiver: ", e.Receiver)
}
