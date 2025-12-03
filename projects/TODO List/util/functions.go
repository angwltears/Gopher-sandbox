package util

import (
	"time"

	"github.com/k0kubun/pp"
)

type Task struct {
	Title         string
	Text          string
	Completed     bool
	CreationTime  time.Time
	CompletedTime time.Time
}

var TaskList = make(map[string]*Task)

func Add(title, text string) {
	if title == "" || text == "" {
		pp.Println("Ошибка ввода!")
	} else {
		pp.Println("Задача добавлена!")
		TaskList[title] = &Task{
			Title:        title,
			Text:         text,
			Completed:    false,
			CreationTime: time.Now(),
		}

	}
}
func List() {
	for _, task := range TaskList {
		if task != nil {
			pp.Println(task)
		}
	}
	pp.Println("Конец")

}
func Del(t string) {
	if TaskList[t] != nil {
		delete(TaskList, t)
	} else {
		pp.Println("Ошибка, такой задачи не существует")
	}

}
func Done(ttl string) {
	if _, exists := TaskList[ttl]; exists {
		TaskList[ttl].Completed = true
		TaskList[ttl].CompletedTime = time.Now()
	} else {
		pp.Println("Ошибка, такой задачи не существует")
	}

}
