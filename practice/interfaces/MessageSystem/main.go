/*
--Вы разрабатываете систему отправки уведомлений
через разные каналы связи:
1.Email
2.SMS
3.Push-Уведомления
--Необходимо разработать модуль отправки уведомлений
который одним методом может принять массив уведомлений
на отправку по различным каналам связи, после чего, собственно
отправить все эти уведомления
--Вместо настоящей отправки уведомлений необходимо использовать
заглушки, на манер заглушек методов оплаты для модуля проведения
оплат из видео-урока
*/
package main

import (
	mmodule "main/MessageModule"
	"main/MessageModule/methods"
)

func main() {
	msg1 := []mmodule.Msg{
		methods.Email{
			Message:  "Hi Tom",
			Receiver: "Tom",
		},
		methods.Push{
			Message:  "Hi Bill",
			Receiver: "Bill",
		},
		methods.SMS{
			Message:  "Hi Jane",
			Receiver: "Jane",
		},
	}
	mmodule.SendAll(msg1)
}
