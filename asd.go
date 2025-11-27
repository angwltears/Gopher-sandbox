package main

import (
	"bufio"

	"fmt"

	"os"

	"strconv"

	"strings"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

var storage = make(map[int]*Task)

var nextTaskID = 1

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println("Введіть команду")

		if ok := scanner.Scan(); !ok {

			fmt.Println("Помилка вводу")

			return

		}

		text := scanner.Text()

		fields := strings.Fields(text)

		cmd := fields[0]

		if cmd == "exit" {

			return

		} else if cmd == "add" {

			args := strings.Join(fields[1:], " ")

			add(args)

		} else if cmd == "list" {

			list()

		} else if cmd == "delete" {

			if len(fields) < 2 {

				fmt.Println("Помилка: введіть ID для видалення")

				continue

			}

			id, err := strconv.Atoi(fields[1])

			if err != nil {

				fmt.Println("Помилка: ID має бути числом.")

				continue

			}

			del(id)

		} else if cmd == "complete" {

			asc := 0

			for i, v := range fields {

				if i > 0 {

					asc, _ = strconv.Atoi(v)

				}

			}

			complete(asc)

		}

	}

}

func add(title string) {

	newTask := &Task{

		ID: nextTaskID,

		Title: title,

		Completed: false,
	}

	storage[newTask.ID] = newTask

	nextTaskID++

}

func list() {

	for k, v := range storage {

		status := "[ ]"

		if storage[k].Completed {

			status = "[x]"

		}

		fmt.Println(status)

		fmt.Println(v.ID)

		fmt.Println(v.Title)

	}

}

func complete(id int) {

	pointer, ok := storage[id]

	if ok {
		pointer.Completed = true

	}

}

func del(id int) {

	delete(storage, id)

}
