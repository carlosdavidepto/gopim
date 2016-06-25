package tasks

import (
	"strings"
)

type DB struct {
	tasks []string
}

func (d *DB) Unpack(tasks string) {
	tmp := strings.Split(tasks, "\n")

	for i := len(tmp) - 1; i >= 0; i-- {
		if tmp[i] != "" {
			break
		}
		tmp = tmp[:i]
	}

	d.tasks = tmp
}

func (d *DB) Pack() string {
	return strings.Join(d.tasks, "\n") + "\n"
}

func (d *DB) Len() int {
	return len(d.tasks)
}

func (d *DB) Add(task string) {
	d.tasks = append(d.tasks, task)
}

func (d *DB) Del(id int) {
	realid := id - 1

	if realid < 0 || realid > len(d.tasks)-1 {
		return
	}

	d.tasks = append(d.tasks[:realid], d.tasks[(realid+1):]...)
}

func (d *DB) Do(id int) {
	realid := id - 1

	if realid < 0 || realid > len(d.tasks)-1 {
		return
	}

	if d.tasks[realid][:2] == "x " {
		return
	}

	d.tasks[realid] = "x " + d.tasks[realid]
}
