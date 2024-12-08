package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var todoList []string

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter command (add/list/delete/exit): ")
		scanner.Scan()
		input := scanner.Text()
		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "add":
			if len(args) > 1 {
				addTask(strings.Join(args[1:], " "))
			} else {
				fmt.Println("Please provide a task to add. Example: add 'Buy milk'")
			}
		case "list":
			listTasks()
		case "delete":
			if len(args) > 1 {
				deleteTask(args[1])
			} else {
				fmt.Println("Please provide the task number to delete. Example: delete 3. To view task numbers, use the 'list' command.")
			}
		case "exit":
			fmt.Println("Exiting...")
			return
		// TODO: Add a "help" command to display the list of available commands
		// TODO: Add 're-rank' command to change the order of tasks
		default:
			fmt.Println("Unknown command. Please use add/list/delete/exit.")
		}
	}
}

func addTask(task string) {
	todoList = append(todoList, task)
	fmt.Println("Added task:", task)
}

func listTasks() {
	if len(todoList) == 0 {
		fmt.Println("No tasks in the list.")
		return
	}
	for i, task := range todoList {
		fmt.Printf("%d: %s\n", i+1, task)
	}
}

func deleteTask(taskNumber string) {
	var index int
	_, err := fmt.Sscanf(taskNumber, "%d", &index)
	if err != nil || index < 1 || index > len(todoList) {
		fmt.Println("Invalid task number.")
		return
	}
	index-- // Convert to zero-based index
	task := todoList[index]
	todoList = append(todoList[:index], todoList[index+1:]...)
	fmt.Println("Deleted task:", task)
}
