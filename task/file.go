package task

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const tasksFile = "tasks.json"

func SaveTasksToFile() error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		log.Printf("Failed to marshal tasks: %v\n", err)
		return err
	}
	err = ioutil.WriteFile(tasksFile, data, 0644)
	if err != nil {
		log.Printf("Failed to write tasks to file: %v\n", err)
		return err
	}
	log.Println("Tasks saved successfully.")
	return nil
}

func LoadTasksFromFile() error {
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
	for _, task := range tasks {
		if task.ID > lastTaskID {
			lastTaskID = task.ID
		}
	}

	return nil
}
