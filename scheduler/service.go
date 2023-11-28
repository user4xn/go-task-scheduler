package scheduler

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

//Create Instance
func NewScheduler() *Scheduler {
	return &Scheduler{
		Database:  NewDatabase(),
		stopChan:  make(chan struct{}),
		waitGroup: sync.WaitGroup{},
	}
}

//Fucntion add task
func (s *Scheduler) AddTask(taskName string, function func(), executionTime int64, interval time.Duration, isRepeated bool) {
	taskID := generateRandomID()

	task := &Task{
		TaskID:        taskID,
		TaskName:      taskName,
		TaskFunction:  function,
		ExecutionTime: executionTime,
		Interval:      interval,
		IsRepeated:    isRepeated,
	}
	s.Database.Lock()
	defer s.Database.Unlock()
	s.Database.Tasks[taskID] = task

	fmt.Printf("Task added - ID: %d, Name: %s, Execution Time: %v\n", taskID, taskName, time.Unix(executionTime, 0))
}

//Fucntion start service scheduler
func (s *Scheduler) Start() {
	go func() {
		for {
			select {
			case <-s.stopChan:
				return
			case now := <-time.After(time.Second):
				s.Database.Lock()
				//Loop from database
				for _, task := range s.Database.Tasks {
					if now.After(time.Unix(task.ExecutionTime, 0)) {
						//Execute
						s.scheduleTask(task)
						if !task.IsRepeated {
							delete(s.Database.Tasks, task.TaskID)
						}
					}
				}
				s.Database.Unlock()
			}
		}
	}()
}

//Stop Service
func (s *Scheduler) Stop() {
	close(s.stopChan)
	s.waitGroup.Wait()
}

//Execution Schedule
func (s *Scheduler) scheduleTask(task *Task) {
	s.waitGroup.Add(1)
	go func() {
		defer s.waitGroup.Done()
		for {
			select {
			case <-s.stopChan:
				return
			case now := <-time.After(time.Until(time.Unix(task.ExecutionTime, 0))):
				fmt.Printf("\nTask running - ID[%d]:\n", task.TaskID)
				defer func() {
					if r := recover(); r != nil {
						log.Println(r)
					}
				}()

				task.TaskFunction()
				if !task.IsRepeated {
					s.removeTask(task.TaskID)
					return
				}
				task.ExecutionTime = now.Add(task.Interval).Unix()
			}
		}
	}()
}

//Remove task by id
func (s *Scheduler) removeTask(taskID int) {
	s.Database.Lock()
	defer s.Database.Unlock()

	delete(s.Database.Tasks, taskID)

	fmt.Printf("Task removed - ID: %d\n", taskID)
}

//Generate random func
func generateRandomID() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10000)
}
