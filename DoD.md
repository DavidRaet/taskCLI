# Definition of Done — Task Tracker CLI
**Project:** Task Tracker CLI
**Language:** Go (Golang)
**Category:** CLI Tools — Project-Based Learning
**Version:** 1.0
**Date:** April 21, 2026

---

## 1. Project Overview

The Task Tracker CLI is a command-line application written in Go that allows users to create, manage, and track tasks directly from the terminal. This document defines the criteria that must be met for the project to be considered complete — covering both technical implementation standards and non-technical learning outcomes.

---

## 2. Non-Technical Requirements

These requirements define the learning goals and project hygiene standards that should be satisfied regardless of implementation details.

### 2.1 Learning Objectives

- Demonstrates a working understanding of Go's basic syntax: variables, types, control flow, and functions.
- Shows practical use of Go's standard library (e.g., `os`, `fmt`, `encoding/json`, `time`).
- Reflects an understanding of Go's idiomatic patterns, including error handling with multiple return values rather than exceptions.
- Shows comfort with Go's package system and project structure conventions.
- Learner can explain every line of code they wrote — no copy-paste without comprehension.

### 2.2 Documentation

- A `README.md` exists at the project root and includes: a brief description of the project, installation instructions, a usage guide with example commands, and a list of all supported flags/subcommands.
- All exported functions, types, and constants have Go-style doc comments (`// FunctionName does...`).
- Complex logic blocks have inline comments explaining the "why," not just the "what."

### 2.3 Project Organization

- The project follows standard Go project layout conventions (e.g., `main.go` at the root or under `cmd/`, data/helper logic separated into packages or files by concern).
- A `go.mod` file exists and accurately reflects the module name and Go version.
- No unnecessary files are committed (e.g., no compiled binaries, no `.DS_Store`, no IDE config files — handled via `.gitignore`).

### 2.4 Version Control

- The project is tracked in Git with meaningful, atomic commit messages (e.g., `feat: add delete task command`, `fix: resolve off-by-one in task ID generation`).
- At minimum, there is one commit per major feature added.
- The repository is hosted on GitHub (or equivalent) and publicly accessible.

---

## 3. Technical Requirements

These requirements define the functional and code-quality standards the application must meet.

### 3.1 Core Functionality

The CLI must support the following operations as subcommands:

| Command | Description |
|---|---|
| `add "<title>"` | Creates a new task with a title, assigns a unique ID, and sets status to `todo`. |
| `update <id> "<title>"` | Updates the title of an existing task by its ID. |
| `delete <id>` | Removes a task permanently by its ID. |
| `mark-in-progress <id>` | Changes a task's status to `in-progress`. |
| `mark-done <id>` | Changes a task's status to `done`. |
| `list` | Lists all tasks. |
| `list todo` | Lists only tasks with status `todo`. |
| `list in-progress` | Lists only tasks with status `in-progress`. |
| `list done` | Lists only tasks with status `done`. |

### 3.2 Data Model

Each task must contain the following fields:

- `id` — a unique integer identifier, auto-incremented.
- `title` — a string describing the task.
- `status` — one of three values: `todo`, `in-progress`, `done`.
- `createdAt` — a timestamp (RFC3339 format) set when the task is created.
- `updatedAt` — a timestamp (RFC3339 format) updated whenever the task is modified.

### 3.3 Persistence

- Tasks are persisted to a local JSON file (e.g., `tasks.json`) in the current working directory.
- The file is created automatically on first run if it does not exist.
- The application reads from this file on every command and writes back after any mutation.
- No external database or third-party storage library is used — standard library only.

### 3.4 CLI Behavior & UX

- The app is invoked via a single compiled binary (e.g., `./task-cli`).
- All subcommands are parsed from `os.Args` directly — no third-party CLI framework (e.g., no `cobra`, no `urfave/cli`). This is intentional for learning purposes.
- Running the binary with no arguments or with an unrecognized command prints a help message listing available commands.
- Invalid IDs (non-existent or non-numeric) produce a clear, user-friendly error message and exit with a non-zero status code.
- Successful operations print a confirmation message to `stdout` (e.g., `Task added successfully (ID: 3)`).
- Errors are printed to `stderr`, not `stdout`.

### 3.5 Error Handling

- All errors are handled explicitly — no blank `_` ignores on error return values unless justified with a comment.
- File I/O errors (read, write, parse) are caught and surfaced with a descriptive message.
- The application never panics under normal usage conditions (missing file, bad input, wrong number of args, etc.).

### 3.6 Code Quality

- Code is formatted with `gofmt` — no formatting violations.
- `go vet ./...` passes with zero warnings.
- There is no dead code, unused imports, or unused variables.
- Logic is organized into focused functions — no single function exceeds ~40–50 lines without clear justification.
- Magic numbers and string literals used more than once are extracted into named constants.

### 3.7 Build & Portability

- The project builds successfully with `go build ./...` on a clean machine with only Go installed.
- The compiled binary runs on the developer's OS without additional dependencies.
- No `CGO` is required (pure Go implementation).

---

## 4. Stretch Goals (Optional but Encouraged)

These are not required for the DoD but are recommended as extensions once the core is complete.

- Add a `due` date field to tasks and a `list overdue` command.
- Add color-coded output using ANSI codes (without third-party libraries) to distinguish task statuses.
- Write at least 3 unit tests using Go's built-in `testing` package — covering task creation, status update, and JSON serialization.
- Support a `--file` flag to specify a custom path for the JSON storage file.
- Add a `search "<keyword>"` command that filters tasks by title substring.

---

## 5. Definition of Done Checklist

Use this checklist before marking the project complete.

**Non-Technical**
- [ ] All learning objectives can be demonstrated and explained verbally
- [ ] `README.md` is complete with description, install steps, and usage examples
- [ ] All exported symbols have doc comments
- [ ] Project structure follows Go conventions
- [ ] `go.mod` is present and correct
- [ ] `.gitignore` excludes binaries and system files
- [ ] Git history has meaningful, atomic commits
- [ ] Repository is public and accessible

**Technical — Functionality**
- [ ] `add` command works and persists the task
- [ ] `update` command modifies title and updates `updatedAt`
- [ ] `delete` command removes the task
- [ ] `mark-in-progress` and `mark-done` commands update status correctly
- [ ] `list` (all), `list todo`, `list in-progress`, `list done` all work correctly
- [ ] Tasks are stored in and loaded from `tasks.json`
- [ ] File is created automatically if missing

**Technical — Quality**
- [ ] No third-party dependencies
- [ ] Args parsed from `os.Args` directly
- [ ] All errors handled explicitly — no unhandled `_` discards
- [ ] App never panics on bad input
- [ ] Errors go to `stderr`, success messages to `stdout`
- [ ] `gofmt` passes with no changes
- [ ] `go vet ./...` passes clean
- [ ] No unused imports or variables
- [ ] `go build ./...` succeeds on a clean machine

---

*This document should be revisited and updated if the project scope changes during development.*