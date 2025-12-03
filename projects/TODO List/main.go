package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todo-list/util"

	"github.com/k0kubun/pp"
)

func main() {
	defer fmt.Println("Программа завершила работу")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		pp.Println("Введите команду")
		if ok := scanner.Scan(); !ok {
			pp.Println("Ошибка ввода")
			util.Logs("", ok)
			continue
		}
		text := scanner.Text()
		if text == "" {
			pp.Println("Вы передали пустую строку")
			util.Logs("", false)
			continue
		}
		util.Logs(text, true)
		fields := strings.Fields(text)
		cmd := fields[0]
		if cmd == "exit" {
			break
		}
		if cmd == "help" {
			pp.Println(
				"help — эта команда позволяет узнать доступные команды и их формат")
			pp.Println(
				"add {заголовок задачи из одного слова} {текст задачи из одного или нескольких слов} — эта команда позволяет добавлять новые задачи в список задач")
			pp.Println(
				"list — эта команда позволяет получить полный список всех задач")
			pp.Println(
				"del {заголовок существующей задачи} — эта команда позволяет удалить задачу по её заголовку")
			pp.Println(
				"done {заголовок существующей задачи} — эта команда позволяет отменить задачу как выполненную")
			pp.Println("events — эта команда позволяет получить список всех событий")
			pp.Println("exit — эта команда позволяет завершить выполнение программы")
		}
		if cmd == "add" || cmd == "Add" || cmd == "ADD" {
			textarg := ""
			for i, arg := range fields {
				if i > 1 {
					textarg += " " + arg
				}
			}
			util.Add(fields[1], textarg)
			continue
		} else if cmd == "done" || cmd == "Done" || cmd == "DONE" {
			if util.TaskList[fields[1]].Completed != true {
				util.Done(fields[1])
				pp.Println("Задача отмечена как выполненная")
			} else {
				pp.Println(" Задача уже выполнена!")
			}
			continue
		} else if cmd == "events" || cmd == "Events" || cmd == "EVENTS" {
			pp.Println("Логи: ")
			for i := 1; i < len(util.Loglist); i++ {
				pp.Println("Строка: ", util.Loglist[i].Str)
				pp.Println("Время: ", util.Loglist[i].Time)
			}
			pp.Println("Конец логов")
			continue
		} else if cmd == "list" || cmd == "List" || cmd == "LIST" {
			util.List()
			continue
		} else if cmd == "del" || cmd == "Del" || cmd == "DEL" {
			util.Del(fields[1])
			pp.Println("Задача успешно удалена!")
			continue
		} else {
			pp.Println("Неизвестная команда")
		}

	}
}
