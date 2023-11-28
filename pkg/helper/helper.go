package helper

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"go-task-scheduler/scheduler"
)

func AddTask(s *scheduler.Scheduler) {
	fmt.Println("\nAdd Task:")
	fmt.Print("Task Name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	taskName := scanner.Text()

	fmt.Print("Execution Time (2006-01-02 15:04:05): ")
	scanner.Scan()
	executionTimeStr := scanner.Text()
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}
	executionTime, err := time.ParseInLocation("2006-01-02 15:04:05", executionTimeStr, location)
	if err != nil {
		fmt.Println("Invalid timestamp. Task not added.")
		return
	}

	fmt.Println(time.Now().Add(5 * time.Second))
	fmt.Println(executionTime)

	fmt.Print("Interval (in seconds, 0 for non-repeated): ")
	scanner.Scan()
	intervalStr := scanner.Text()
	interval, err := strconv.Atoi(intervalStr)
	if err != nil {
		fmt.Println("Invalid interval. Task not added.")
		return
	}

	fmt.Print("Is Repeated (true/false): ")
	scanner.Scan()
	isRepeatedStr := scanner.Text()
	isRepeated, err := strconv.ParseBool(isRepeatedStr)
	if err != nil {
		fmt.Println("Invalid input for 'Is Repeated'. Task not added.")
		return
	}

	s.AddTask(taskName, func() {
		fmt.Printf("Executing %s at %v\n", taskName, time.Now())
	}, executionTime.Unix(), time.Duration(interval)*time.Second, isRepeated)

	fmt.Printf("Task '%s' added.\n", taskName)
}

func ListTasks(s *scheduler.Scheduler) {
	fmt.Println("\nList of Tasks:")
	for _, task := range s.Database.Tasks {
		fmt.Printf("Task ID: %d, Task Name: %s, Execution Time: %v\n", task.TaskID, task.TaskName, time.Unix(task.ExecutionTime, 0))
	}
}

func RemoveTask(s *scheduler.Scheduler) {
	fmt.Print("\nEnter Task ID to remove: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	taskIDStr := scanner.Text()

	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		fmt.Println("Invalid Task ID. Task not removed.")
		return
	}

	s.Database.Lock()
	defer s.Database.Unlock()

	if _, ok := s.Database.Tasks[taskID]; ok {
		delete(s.Database.Tasks, taskID)
		fmt.Printf("Task with ID %d removed.\n", taskID)
	} else {
		fmt.Printf("Task with ID %d not found. Task not removed.\n", taskID)
	}
}
