package methods

import "fmt"

type SMS struct {
	Message  string
	Receiver string
}

func (e SMS) Send() {
	fmt.Println("Sms message:", e.Message)
	fmt.Println("SMs receiver: ", e.Receiver)
}
