package tasks

import (
  "testing"
)

func TestEmptyStringHasNoTasks(t *testing.T) {
  db := &DB{}
  db.Unpack("")
  if db.Len() != 0 {
    t.Errorf("empty database should not contain tasks")
  }
}

func TestOneLineIsOneTask(t * testing.T) {
  db := &DB{}
  db.Unpack("one simple task")
  if db.Len() != 1 {
    t.Errorf("one line should correspond to one task")
  }
}

func TestEmptyLinesAtEndIgnored(t *testing.T) {
  db := &DB{}
  db.Unpack("task\n\n\n")
  if db.Len() != 1 {
    t.Errorf("empty lines at the end should be ignored")
  }
}

func TestEmptyLinesInTheMiddleAreKept(t *testing.T) {
  db := &DB{}
  db.Unpack("t\n\nt\n")
  if db.Len() != 3 {
    t.Errorf("empty lines in the middle should be ignored")
  }
}

func TestPackNoModification(t *testing.T) {
  db := &DB{}
  db.Unpack("task1\n")
  if db.Pack() != "task1\n" {
    t.Errorf("Pack() modifies db in unexpected way")
  }
}

func TestAddTask(t *testing.T) {
  db := &DB{}
  db.Unpack("task1\n")
  db.Add("task2")
  if db.Pack() != "task1\ntask2\n" {
    t.Errorf("Pack() modifies db in unexpected way")
  } 
}
