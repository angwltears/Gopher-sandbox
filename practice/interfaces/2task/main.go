package main

import "fmt"

type Notifier interface {
	Send(string)
}
type Email struct {
	Adress string
}
type SMS struct {
	Phone string
}

func (s SMS) Send(msg string) {
	fmt.Println("Отправка сообщения на телефон: ", s.Phone, "сообщение: ", msg)

}
func (E Email) Send(msg string) {
	fmt.Println("Отправка сообщения на адрес: ", E.Adress, "сообщение: ", msg)

}
func NotifyUser(n Notifier, msg string) {
	n.Send(msg)
}
func main() {
	e1 := Email{
		Adress: "you",
	}
	e2 := SMS{
		Phone: "123",
	}
	NotifyUser(e1, "12345")
	NotifyUser(e2, "123456")
}
