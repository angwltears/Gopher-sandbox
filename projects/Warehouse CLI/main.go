package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var nextID int = 1

type Product struct {
	ID       int
	Name     string
	Price    int
	Quantity int
}

var Storage = make(map[int]*Product)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Введите команду: ")
		ok := scanner.Scan()
		if !ok {
			fmt.Println("Ошибка ввода ")
			return
		}
		text := scanner.Text()
		fields := strings.Fields(text)
		cmd := fields[0]
		if cmd == "add" {
			price, err := strconv.ParseInt(fields[2], 10, 64)
			qty, err2 := strconv.ParseInt(fields[3], 10, 64)
			if err == nil && err2 == nil {
				add(fields[1], int(price), int(qty))
			} else {
				fmt.Println("Неверно переданы аргументы")
			}

		} else if cmd == "list" {
			list()
		} else if cmd == "stock" {
			id, err := strconv.ParseInt(fields[1], 10, 64)
			amount, err2 := strconv.ParseInt(fields[2], 10, 64)
			if err == nil && err2 == nil {
				stock(int(id), int(amount))

			} else {
				return
			}

		} else if cmd == "sell" {
			id, err := strconv.ParseInt(fields[1], 10, 64)
			if err == nil {
				sell(int(id))
			}
		}
	}
}
func add(name string, price, qty int) {
	newProduct := &Product{
		ID:       nextID,
		Name:     name,
		Price:    price,
		Quantity: qty,
	}
	Storage[nextID] = newProduct
	nextID++
}
func stock(id, amount int) {
	if Storage[id] != nil {
		prod := Storage[id]
		if prod.Quantity+amount < 0 {
			fmt.Println("Ошибка, amount не может быть меньше 0...")
		} else {
			prod.Quantity += amount
		}
	} else {
		fmt.Println("Ошибка: ID не существует")
	}

}
func list() {
	total := 0
	for k, v := range Storage {
		fmt.Printf("%d\t%s\t%d\t%d\n", k, v.Name, v.Quantity, v.Price)
		total += v.Price * v.Quantity

	}
	fmt.Println("тотал: ", total)
}
func sell(id int) {
	if Storage[id] != nil {
		if Storage[id].Quantity > 0 {
			Storage[id].Quantity--
			fmt.Println("Продано: ", 1, "Осталось: ", Storage[id].Quantity)
		} else {
			fmt.Println("Ошибка: указанное значение превышает наличие, разница: ")
		}

	} else {
		fmt.Println("Ошибка ID")
	}
}
