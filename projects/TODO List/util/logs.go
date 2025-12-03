package util

import "time"

var index int = 1

type loginfo struct {
	Str  string
	Time time.Time
}

var Loglist = make(map[int]*loginfo)

func Logs(arg string, ok bool) {
	if ok && arg != "" {
		Loglist[index] = &loginfo{Str: arg, Time: time.Now()}
	} else {
		Loglist[index] = &loginfo{Str: "Input error", Time: time.Now()}
	}
	index++

}
