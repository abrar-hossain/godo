# godo - A CLI Todo App in Go

`godo` is a command-line todo list manager written in Go.  
It provides the ability to add, view, complete, and delete tasks directly from the terminal.

Tasks are saved in a local JSON file, ensuring the list persists even after closing the program.

---

## Features

- Add tasks with a single command
- List all tasks with status indicators (done or pending)
- Mark tasks as done
- Delete specific tasks
- Clear the entire task list
- Unit tested using Go's testing framework
- Data stored locally in `todos.json`

---

## How It Works

Each time the program is run, it performs the following steps:

1. Parses the command (`add`, `list`, `done`, `delete`, etc.)
2. Loads the task list from `todos.json`
3. Applies the requested action
4. Saves the updated list back to the file

---

## Flowchart

This diagram illustrates the logic of the application:

![CLI Todo Flowchart](flowchart.png)

---

## What I Learned

This project provided practical experience with core concepts in Go, including:

- Memory management through slices and byte arrays
- Command-line argument parsing using `os.Args`
- JSON encoding and decoding using `encoding/json`
- Organizing and modularizing a Go project
- Writing unit tests using table-driven testing
- Understanding Go's performance model and why it is fast
- Building tools from scratch using only the standard library

The experience helped solidify understanding of Go's simplicity, power, and efficiency.

---


