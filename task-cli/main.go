package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

const fileName = "tasks.json"

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command> [arguments]")
		return
	}

	command := os.Args[1]

	switch command {

	case "add":
		addTask(os.Args[2:])

	case "update":
		updateTask(os.Args[2:])

	case "delete":
		deleteTask(os.Args[2:])

	case "mark-in-progress":
		updateStatus(os.Args[2:], "in-progress")

	case "mark-done":
		updateStatus(os.Args[2:], "done")

	case "list":
		listTasks(os.Args[2:])

	default:
		fmt.Println("Unknown command")
	}
}



func loadTasks() []Task {

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		os.WriteFile(fileName, []byte("[]"), 0644)
	}

	data, _ := os.ReadFile(fileName)

	var tasks []Task
	json.Unmarshal(data, &tasks)

	return tasks
}


func saveTasks(tasks []Task) {

	jsonData, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile(fileName, jsonData, 0644)
}


func getNextID(tasks []Task) int {

	maxID := 0

	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	return maxID + 1
}

func addTask(args []string) {

	if len(args) < 1 {
		fmt.Println("Usage: add \"task description\"")
		return
	}

	description := args[0]

	tasks := loadTasks()

	newTask := Task{
		ID:          getNextID(tasks),
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)

	saveTasks(tasks)

	fmt.Println("Task added successfully (ID:", newTask.ID, ")")
}


func updateTask(args []string) {

	if len(args) < 2 {
		fmt.Println("Usage: update <id> \"new description\"")
		return
	}

	id, _ := strconv.Atoi(args[0])
	newDesc := args[1]

	tasks := loadTasks()

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = newDesc
			tasks[i].UpdatedAt = time.Now()

			saveTasks(tasks)
			fmt.Println("Task updated successfully")
			return
		}
	}

	fmt.Println("Task not found")
}


func deleteTask(args []string) {

	if len(args) < 1 {
		fmt.Println("Usage: delete <id>")
		return
	}

	id, _ := strconv.Atoi(args[0])

	tasks := loadTasks()

	for i := range tasks {
		if tasks[i].ID == id {

			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks(tasks)

			fmt.Println("Task deleted successfully")
			return
		}
	}

	fmt.Println("Task not found")
}


func updateStatus(args []string, status string) {

	if len(args) < 1 {
		fmt.Println("Usage: mark-xxx <id>")
		return
	}

	id, _ := strconv.Atoi(args[0])

	tasks := loadTasks()

	for i := range tasks {
		if tasks[i].ID == id {

			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()

			saveTasks(tasks)

			fmt.Println("Task status updated")
			return
		}
	}

	fmt.Println("Task not found")
}

func listTasks(args []string) {

	tasks := loadTasks()

	filter := ""

	if len(args) > 0 {
		filter = args[0]
	}

	for _, task := range tasks {

		if filter != "" && task.Status != filter {
			continue
		}

		fmt.Printf("[%d] %s (%s)\n",
			task.ID,
			task.Description,
			task.Status,
		)
	}
}
