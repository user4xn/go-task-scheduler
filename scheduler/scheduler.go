package scheduler

import (
	"sync"
	"time"
)

type Task struct {
	TaskID        int
	TaskName      string
	TaskFunction  func()
	ExecutionTime int64
	Interval      time.Duration
	IsRepeated    bool
}

type Scheduler struct {
	Database  *Database
	stopChan  chan struct{}
	waitGroup sync.WaitGroup
}
