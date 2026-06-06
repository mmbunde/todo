# Todo CLI

A command-line task manager built in Go, designed as a learning project that grows in complexity through multiple phases. Each phase introduces new Go concepts, starting simple and progressively building toward a production-grade application.

---

## The Journey

This project isn't just a task manager — it's a structured way to learn Go by building something real. Rather than following tutorials, each phase solves a genuine problem and introduces concepts naturally as they're needed.

---

## Phases

### ✅ Phase 1 — Basic CLI Task Manager
**Goal:** Create a simple terminal application where tasks exist only while the app is running.

**What I learned:**
- Structs and how Go models data
- Slices and dynamic collections
- Functions, parameters, and return values
- Reading command-line input
- Go's explicit error handling philosophy

---

### ✅ Phase 2 — Persistent Storage (JSON)
**Goal:** Make tasks survive after closing the app by saving to a JSON file.

**What I learned:**
- File I/O with `os.ReadFile` and `os.WriteFile`
- JSON serialization with `json.Marshal` and `json.Unmarshal`
- Startup initialization and handling missing files
- Separating concerns into multiple files

---

### ✅ Phase 3 — Professional CLI Design
**Goal:** Turn the toy CLI into a professional developer tool using Cobra.

**What I learned:**
- CLI frameworks and enterprise-grade patterns with [Cobra](https://github.com/spf13/cobra)
- Subcommands, persistent flags, and auto-generated help menus
- Package organization (`/cmd`, `/models`, `/storage`)
- Shell autocompletion
- `PersistentPreRun` for shared validation logic

---

### ✅ Phase 4 — Database Integration (SQLite)
**Goal:** Replace JSON file storage with SQLite for proper data persistence.

**What I learned:**
- SQL basics — `SELECT`, `INSERT`, `UPDATE`, `DELETE`
- Using Go's `database/sql` package
- Prepared statements and SQL injection prevention
- Handling database-specific errors
- Repository pattern — separating storage logic from business logic

---

### 🚧 Phase 5 — REST API *(planned)*
**Goal:** Turn the app into a network service.

**Planned features:**
- HTTP endpoints for creating, fetching, and updating tasks
- JSON request/response handling
- Middleware for logging, authentication, and request timing
- Using `net/http` or Gin

---

### 💡 Future Ideas
- Project/list support — separate task lists per project
- UUID-based task IDs
- Priority levels and due dates
- Filtering by status or priority

---

## Quick Start

### Installation

```bash
git clone https://github.com/mmbunde/todo
cd todo
go build -o todo
sudo mv todo /usr/local/bin/
```

### Shell Completion (bash)

```bash
source <(todo completion bash)
```

To make it permanent, add the above line to your `~/.bashrc`.

---

## Usage

```bash
todo -f <database.db> <command> [task title]
```

### Commands

| Command | Description | Example |
|---|---|---|
| `add` | Add a new task | `todo -f tasks.db add "Learn Go"` |
| `list` | List all tasks | `todo -f tasks.db list` |
| `complete` | Mark a task complete | `todo -f tasks.db complete "Learn Go"` |
| `delete` | Delete a task | `todo -f tasks.db delete "Learn Go"` |
| `completion` | Generate shell completion script | `todo completion bash` |

### Flags

| Flag | Short | Description |
|---|---|---|
| `--file` | `-f` | SQLite database file to use |
| `--help` | `-h` | Show help for any command |

### Examples

```bash
# Add some tasks
todo -f work.db add "Review pull requests"
todo -f work.db add "Write tests"

# List all tasks
todo -f work.db list

# Mark a task complete
todo -f work.db complete "Write tests"

# Delete a task
todo -f work.db delete "Review pull requests"
```

---

## Project Structure

```
todo/
├── cmd/
│   ├── root.go        # Root command, persistent flags, and DB initialization
│   ├── add.go         # Add a task
│   ├── list.go        # List all tasks
│   ├── complete.go    # Mark a task complete
│   ├── delete.go      # Delete a task
│   ├── completion.go  # Shell completion
│   └── helpers.go     # Shared utility functions
├── storage/
│   └── storage.go     # Database initialization and file path helpers
├── main.go
├── go.mod
└── go.sum
```

---
