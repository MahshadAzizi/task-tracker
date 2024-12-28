package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

var tasks = make([]Task, 0)

const tasksFile = "tasks.json"

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func main() {
	fmt.Println("Welcome to the Task Tracker project")
	err := loadTasksFromFile()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}
	command, args := getUserInput()
	switch command {
	case "add":
		if len(args) == 0 {
			fmt.Println("Please provide a description")
		} else {
			task := addTask(args[0])
			tasks = append(tasks, task)
			fmt.Printf("Task added successfully (ID: %d)\n", task.ID)
		}
	case "update":
		if len(args) == 0 {
			fmt.Println("Please provide a description")
			return
		} else {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Error during conversion")
				return
			}
			updateTask(id, args[1])
		}
	case "list":
		if len(args) == 0 {
			fmt.Printf("Tasks %v\n", tasks)
		} else {
			var finalList []Task
			for _, task := range tasks {
				if task.Status == args[0] {
					finalList = append(finalList, task)
				}
			}
			fmt.Println(finalList)

		}

	case "delete":
		if len(args) == 0 {
			fmt.Println("Please provide a description")
		} else {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Error during conversion")
			}
			deleteTask(id)
		}
	case "mark-in-progress":
		if len(args) == 0 {
			fmt.Println("Please provide a description")
		} else {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Error during conversion")
			}
			updateTaskStatus(id, "in-progress")
		}
	case "mark-done":
		if len(args) == 0 {
			fmt.Println("Please provide a description")
		} else {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Error during conversion")
			}
			updateTaskStatus(id, "done")
		}
	}
}

func getUserInput() (string, []string) {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a command with arguments (split by spaces):")
	reader.Scan()
	input := reader.Text()
	parts := strings.Fields(input)
	command := parts[0]
	args := parts[1:]
	return command, args
}

func saveTasksToFile() error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal tasks: %v", err)
	}

	err = ioutil.WriteFile(tasksFile, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write tasks to file: %v", err)
	}

	return nil
}

func loadTasksFromFile() error {
	if _, err := os.Stat(tasksFile); os.IsNotExist(err) {
		return nil
	}

	data, err := ioutil.ReadFile(tasksFile)
	if err != nil {
		return fmt.Errorf("failed to read tasks file: %v", err)
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return fmt.Errorf("failed to unmarshal tasks: %v", err)
	}

	return nil
}

func addTask(description string) Task {
	task := Task{
		ID:          len(tasks) + 1,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append(tasks, task)

	err := saveTasksToFile()
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
	}
	return task
}

func deleteTask(id int) {
	for index, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:index], tasks[index+1:]...)

			err := saveTasksToFile()
			if err != nil {
				fmt.Printf("Error saving tasks: %v\n", err)
			}
		}
	}
}

func updateTask(id int, description string) {
	for index, task := range tasks {
		if task.ID == id {
			tasks[index].Description = description
			tasks[index].UpdatedAt = time.Now()
			err := saveTasksToFile()
			if err != nil {
				fmt.Printf("Error saving tasks: %v\n", err)
			}
		}
	}
}

func updateTaskStatus(id int, status string) {
	for index, task := range tasks {
		if task.ID == id {
			tasks[index].Status = status
			tasks[index].UpdatedAt = time.Now()
			err := saveTasksToFile()
			if err != nil {
				fmt.Printf("Error saving tasks: %v\n", err)
			}
		}
	}
}
