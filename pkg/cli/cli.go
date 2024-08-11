package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hiden2000/taskmaster/internal/storage"
	"github.com/hiden2000/taskmaster/internal/task"
)

type CLI struct {
	storage *storage.Storage
}

func NewCLI(s *storage.Storage) *CLI {
	return &CLI{storage: s}
}

func (c *CLI) Run() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("taskmaster> ")
		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, " ")

		if input == "exit" {
			break
		}

		c.executeCommand(input)
	}
}

func (c *CLI) executeCommand(input string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return
	}

	command := parts[0]
	args := parts[1:]

	switch command {
	case "add":
		c.addTask(args)
	case "list":
		c.listTasks()
	case "complete":
		c.completeTask(args)
	case "delete":
		c.deleteTask(args)
	default:
		fmt.Println("Unknown command. Available commands: add, list, complete, delete")
	}
}

func (c *CLI) addTask(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: add <title> <descrption>")
		return
	}

	title := args[0]
	description := strings.Join(args[1:], " ")
	newTask := task.NewTask(0, title, description, time.Now().Add(24*time.Hour))

	err := c.storage.AddTask(newTask)
	if err != nil {
		fmt.Printf("Error adding task: %v\n", err)
		return
	}

	fmt.Printf("Task added: %s\n", newTask)
}

func (c *CLI) listTasks() {
	tasks := c.storage.ListTasks()
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
	}

	for _, t := range tasks {
		fmt.Println(t)
	}
}

func (c *CLI) completeTask(args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: complete <task_id>")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Error: <task_id> must be integer, but %s given.", args[0])
		return
	}

	task, err := c.storage.GetTask(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	task.Complete()
	err = c.storage.UpdateTask(task)
	if err != nil {
		fmt.Printf("Error updating task: %v\n", err)
		return
	}

	fmt.Printf("Task completed: %s\n", task)
}

func (c *CLI) deleteTask(args []string) {

	if len(args) != 1 {
		fmt.Println("Usage: delete <task_id>")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Error: <task_id> must be integer, but %s given.", args[0])
		return
	}
	err = c.storage.DeleteTask(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Task %d deleted\n", id)
}
