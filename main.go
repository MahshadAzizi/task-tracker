package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"task-tracker/task"
)

func main() {
	fmt.Println("Welcome to the Task Tracker Project!")

	if err := task.LoadTasksFromFile(); err != nil {
		log.Printf("Error loading tasks: %v\n", err)
		return
	}

	defer func() {
		if err := task.SaveTasksToFile(); err != nil {
			log.Printf("Error saving tasks: %v\n", err)
		}
	}()

	commands := map[string]func([]string){
		"add":              task.HandleAdd,
		"update":           task.HandleUpdate,
		"list":             task.HandleList,
		"delete":           task.HandleDelete,
		"mark-in-progress": task.HandleMarkInProgress,
		"mark-done":        task.HandleMarkDone,
		"help":             task.HandleHelp,
	}

	for {
		fmt.Println("Enter a command with arguments (type 'help' for list of commands):")
		command, args := getUserInput()
		if command == "" {
			fmt.Println("No command entered. Please try again.")
			continue
		}
		if handler, exists := commands[command]; exists {
			handler(args)
		} else {
			fmt.Printf("Unknown command: %s\n", command)
		}
	}
}

func getUserInput() (string, []string) {
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	input := reader.Text()
	parts := strings.Fields(input)
	if len(parts) == 0 {
		fmt.Println("No command entered. Please try again.")
		return "", []string{}
	}
	command := parts[0]
	args := parts[1:]
	return command, args
}
