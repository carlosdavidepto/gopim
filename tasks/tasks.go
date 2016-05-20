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
    if (tmp[i] != "") {
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
