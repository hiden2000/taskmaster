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

func (c *CLI) Run() error {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("taskmaster> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("error reading input: %w", err)
		}
		input = strings.Trim(input, " ")

		if input == "exit" {
			break
		}

		if err := c.executeCommand(input); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
	}
	return nil
}

func (c *CLI) executeCommand(input string) error {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil
	}

	command := parts[0]
	args := parts[1:]

	switch command {
	case "add":
		return c.addTask(args)
	case "list":
		return c.listTasks()
	case "complete":
		return c.completeTask(args)
	case "delete":
		return c.deleteTask(args)
	default:
		return fmt.Errorf("unknown command. Available commands: add, list, complete, delete")
	}
}

func (c *CLI) addTask(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: add <title> <descrption>")
	}

	title := args[0]
	description := strings.Join(args[1:], " ")
	newTask := task.NewTask(0, title, description, time.Now().Add(24*time.Hour))

	err := c.storage.AddTask(newTask)
	if err != nil {
		return fmt.Errorf("error adding task: %v\n", err)
	}

	fmt.Printf("Task added: %s\n", newTask)
	return nil
}

func (c *CLI) listTasks() error {
	tasks := c.storage.ListTasks()
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}

	for _, t := range tasks {
		fmt.Println(t)
	}
	return nil
}

func (c *CLI) completeTask(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: complete <task_id>")

	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("error: <task_id> must be integer, but %s given.", args[0])
	}

	task, err := c.storage.GetTask(id)
	if err != nil {
		return fmt.Errorf("error getting task: %w", err)
	}

	task.Complete()
	err = c.storage.UpdateTask(task)
	if err != nil {
		return fmt.Errorf("error updating task: %w\n", err)
	}

	fmt.Printf("Task completed: %s\n", task)
	return nil
}

func (c *CLI) deleteTask(args []string) error {

	if len(args) != 1 {
		return fmt.Errorf("usage: delete <task_id>")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("error: <task_id> must be integer, but %s given.", args[0])
	}
	err = c.storage.DeleteTask(id)
	if err != nil {
		return fmt.Errorf("error deleting task: %w", err)

	}

	fmt.Printf("Task %d deleted\n", id)
	return nil
}
