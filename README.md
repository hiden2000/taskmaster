# Taskmaster

Taskmaster is a simple command-line task management application written in Go. It allows users to add, list, complete, and delete tasks through an interactive CLI interface.

## Features

- Add new tasks with a title and description
- List all tasks
- Mark tasks as complete
- Delete tasks
- Simple in-memory storage (data is not persistent between runs)

## Installation

To install Taskmaster, make sure you have Go installed on your system, then run:

```bash
go get github.com/hiden2000/taskmaster
```

Alternatively, you can clone the repository and build it manually:

```bash
git clone https://github.com/hiden2000/taskmaster.git
cd taskmaster
go build -o taskmaster cmd/taskmaster/main.go
```

## Usage

To start Taskmaster, run the following command in your terminal:

```bash
./taskmaster
```

Once the application is running, you can use the following commands:

- `add <title> <description>`: Add a new task
- `list`: List all tasks
- `complete <task_id>`: Mark a task as complete
- `delete <task_id>`: Delete a task
- `exit`: Exit the application

## Example

```
taskmaster> add Buy groceries Pick up milk and eggs
Task added: Task 1: Buy groceries [Pending] (Due: 2023-05-15)

taskmaster> list
Task 1: Buy groceries [Pending] (Due: 2023-05-15)

taskmaster> complete 1
Task completed: Task 1: Buy groceries [Completed] (Due: 2023-05-15)

taskmaster> delete 1
Task 1 deleted

taskmaster> exit
```

## Project Structure

- `cmd/taskmaster/main.go`: Entry point of the application
- `internal/storage/storage.go`: In-memory storage implementation
- `internal/task/task.go`: Task struct and related methods
- `pkg/cli/cli.go`: CLI interface and command processing

## Future Improvements

- Persistent storage (e.g., file-based or database)
- Task priority levels
- Due date management
- Improved error handling and user feedback
- Configuration options (e.g., default due date)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.