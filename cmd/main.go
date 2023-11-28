package main

import (
	"bufio"
	"fmt"
	"os"

	"go-task-scheduler/pkg/helper"
	"go-task-scheduler/scheduler"
)

func main() {
	// Create a new scheduler
	s := scheduler.NewScheduler()

	// Run the CLI loop
	cliLoop(s)

	// Stop the scheduler when the CLI loop exits
	s.Stop()
}

func cliLoop(s *scheduler.Scheduler) {
	fmt.Println("\n\nTask Scheduler CLI")
	fmt.Println("===================")

	for {
		// Display CLI options
		fmt.Println("Options:")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Remove Task")
		fmt.Println("4. Run Schedule")
		fmt.Println("5. Exit")

		// Prompt user for option
		fmt.Print("Select option (1-4):")

		// Read user input
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		option := scanner.Text()

		// Process user option
		switch option {
		case "1":
			// Option to add a new task
			helper.AddTask(s)
		case "2":
			// Option to list tasks
			helper.ListTasks(s)
		case "3":
			// Option to remove a task
			helper.RemoveTask(s)
		case "4":
			// Option to start the scheduler in a goroutine
			go s.Start()
			fmt.Println("Schedule Running...")
		case "5":
			// Option to exit the CLI
			fmt.Println("Exiting CLI.")
			return
		default:
			// Invalid option
			fmt.Println("Invalid option. Please enter a number between 1 and 5.")
		}

		// Add a newline for better readability
		fmt.Println("")
	}
}
