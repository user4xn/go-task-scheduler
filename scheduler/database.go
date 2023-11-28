package scheduler

import "sync"

type Database struct {
	Tasks map[int]*Task
	sync.Mutex
}

func NewDatabase() *Database {
	return &Database{
		Tasks: make(map[int]*Task),
	}
}
