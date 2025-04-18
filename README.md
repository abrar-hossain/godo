# godo - A CLI Todo App in Go 🐹✅

`godo` is a command-line todo list manager written in Go.  
It helps you stay organized by allowing you to **add, view, complete, and delete tasks** right from your terminal.

Tasks are saved in a local JSON file, so your list persists even after you quit.

---

## ✨ Features

- ➕ Add tasks with a single command
- 📋 List all tasks with status indicators
- ✅ Mark tasks as done
- ❌ Delete completed or unwanted tasks
- 🧹 Clear the entire list
- 🧪 Unit tested with `go test`
- 💾 Tasks stored locally in `todos.json`

---

## ⚙️ How It Works

Each time you run `godo`, it:

1. Parses your command (`add`, `list`, `done`, `delete`, etc.)
2. Loads tasks from `todos.json`
3. Applies the action
4. Saves the updated list back to the file

---

## 🧭 Flowchart

This visual explains the logic of the app:

![CLI Todo Flowchart](flowchart.png)

---

