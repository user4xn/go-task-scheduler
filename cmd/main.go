package main

import (
	"bufio"
	"fmt"
	"os"

	"go-task-scheduler/scheduler"
	"go-task-scheduler/pkg/helper"
)

func main() {
	s := scheduler.NewScheduler()

	cliLoop(s)

	s.Stop()
}

func cliLoop(s *scheduler.Scheduler) {
	fmt.Println("\n\nTask Scheduler CLI")
	fmt.Println("===================")

	for {
		fmt.Println("Options:")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Remove Task")
		fmt.Println("4. Run Schedule")
		fmt.Println("5. Exit")

		fmt.Print("Select option (1-4):")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			helper.AddTask(s)
		case "2":
			helper.ListTasks(s)
		case "3":
			helper.RemoveTask(s)
		case "4":
			go s.Start()
			fmt.Println("Schedule Running...")
		case "5":
			fmt.Println("Exiting CLI.")
			return
		default:
			fmt.Println("Invalid option. Please enter a number between 1 and 4.")
		}

		fmt.Println("")
	}
}
