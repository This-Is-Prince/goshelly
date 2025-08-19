# Minimal Shell in Golang

## Overview

This project is a minimal shell written in Go, designed to provide basic shell functionality. It supports built-in commands and the ability to execute external commands.

## Features

- **Built-in Commands**:

  - `cd`: Change the current directory.
  - `echo`: Print arguments to the standard output.
  - `exit`: Exit the shell with an optional exit code.
  - `pwd`: Print the current working directory.
  - `type`: Display information about a command (whether it's built-in or external).
- **External Commands**:

  - Executes commands available in the system's `PATH`.
- **Interactive REPL**:

  - A Read-Eval-Print Loop (REPL) for continuous command execution.

## Project Structure

- `app/main.go`: Entry point of the application.
- `app/builtin/`: Contains implementations of built-in commands.
  - `cd.go`: Implements the `cd` command.
  - `echo.go`: Implements the `echo` command.
  - `exit.go`: Implements the `exit` command.
  - `pwd.go`: Implements the `pwd` command.
  - `type.go`: Implements the `type` command.
- `app/cmd/`: Handles command parsing and execution.
  - `builtin.go`: Maps built-in commands to their implementations.
  - `cmd.go`: Defines the `Cmd` struct and its methods for command execution.
- `app/shell/`: Implements the shell's REPL.
  - `shell.go`: Defines the `Shell` struct and its methods for reading, evaluating, and printing commands.

## How to Run

1. Clone the repository:
   ```bash
   git clone https://github.com/This-Is-Prince/goshelly.git
   cd goshelly
   ```
