package task

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

const (
	StatusTodo       = "todo"
	StatusInProgress = "in-progress"
	StatusDone       = "done"
)

var tasks = make([]Task, 0)
var lastTaskID int

func HandleAdd(args []string) {
	if len(args) < 1 {
		fmt.Println("Error: Missing description for 'add' command")
		return
	}
	description := strings.Join(args, " ")
	task := AddTask(description)
	fmt.Printf("Task added successfully (ID: %d)\n", task.ID)
}

func HandleUpdate(args []string) {
	if len(args) < 2 {
		fmt.Println("Error: Missing arguments for 'update' command (usage: update <id> <description>)")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Error: Invalid task ID")
		return
	}

	description := strings.Join(args[1:], " ")
	if UpdateTask(id, description) {
		fmt.Printf("Task with ID %d updated successfully.\n", id)
	} else {
		fmt.Printf("Task with ID %d not found.\n", id)
	}
}

func HandleList(args []string) {
	status := ""
	if len(args) > 0 {
		status = args[0]
	}

	fmt.Println("Listing tasks:")
	for _, task := range tasks {
		if status == "" || task.Status == status {
			printTask(task)
		}
	}
}

func HandleDelete(args []string) {
	if len(args) < 1 {
		fmt.Println("Error: Missing task ID for 'delete' command")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Error: Invalid task ID")
		return
	}

	if DeleteTask(id) {
		fmt.Printf("Task with ID %d deleted successfully.\n", id)
	} else {
		fmt.Printf("Task with ID %d not found.\n", id)
	}
}
func HandleMarkInProgress(args []string) {
	handleUpdateStatus(args, StatusInProgress)
}

func HandleMarkDone(args []string) {
	handleUpdateStatus(args, StatusDone)
}

func UpdateTask(id int, description string) bool {
	task, index := findTaskByID(id)
	if task == nil {
		return false
	}
	tasks[index].Description = description
	tasks[index].UpdatedAt = time.Now()
	return true
}

func AddTask(description string) Task {
	lastTaskID++
	task := Task{
		ID:          lastTaskID,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append(tasks, task)
	return task
}

func DeleteTask(id int) bool {
	_, index := findTaskByID(id)
	if index == -1 {
		return false
	}
	tasks = append(tasks[:index], tasks[index+1:]...)
	return true
}

func handleUpdateStatus(args []string, status string) {
	if len(args) < 1 {
		fmt.Println("Error: Missing task ID")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Error: Invalid task ID")
		return
	}

	task, index := findTaskByID(id)
	if task == nil {
		fmt.Printf("Task with ID %d not found.\n", id)
		return
	}

	tasks[index].Status = status
	tasks[index].UpdatedAt = time.Now()
	fmt.Printf("Task with ID %d marked as %s.\n", id, status)
}
func findTaskByID(id int) (*Task, int) {
	for index, task := range tasks {
		if task.ID == id {
			return &tasks[index], index
		}
	}
	return nil, -1
}

func printTask(task Task) {
	fmt.Printf("ID: %d | Description: %s | Status: %s | Created: %s | Updated: %s\n",
		task.ID, task.Description, task.Status,
		task.CreatedAt.Format("2006-01-02 15:04:05"),
		task.UpdatedAt.Format("2006-01-02 15:04:05"))
}
