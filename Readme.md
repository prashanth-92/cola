# cola

`cola` is a simple shell implementation written in Go. It provides a lightweight command-line interface with basic file system execution and management capabilities.

## Introduction

`cola` was designed to mimic core functionalities of a standard Unix shell. It handles standard input/output operations, supporting a variety of built-in commands for file manipulation and system navigation.

## Features

- **Interactive Prompt**: A continuous read-eval-print loop (REPL) that accepts user commands.
- **File System Navigation**: Change directories and list contents.
- **File Management**: Create, copy, move, and remove files and directories.
- **Content Viewing**: Display file contents, including head and tail support.
- **Basic Utilities**: Echo text and display current user.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.16 or higher recommended)

### Installation

Clone the repository and navigate to the project directory:

```bash
git clone https://github.com/prashanth/cola.git
cd cola
```

### Running the Shell

You can run the shell directly using `go run`:

```bash
go run main.go
```

Or build an executable:

```bash
go build -o cola
./cola
```

## Usage

Once started, `cola` displays a prompt `> `. You can enter commands as you would in a typical shell environment.

**Example Session:**

```text
> pwd
/Users/prashanth/git/cola
> mkdir test_dir
> cd test_dir
> touch hello.txt
> echo Hello from cola!
Hello from cola!
> exit
```

## Supported Commands

| Command | Description | Usage |
|---------|-------------|-------|
| `ls`    | List directory contents | `ls [directory]` |
| `cd`    | Change the current directory | `cd <directory>` |
| `pwd`   | Print working directory | `pwd` |
| `cat`   | Display file contents | `cat <filename>` |
| `head`  | Display the first 10 lines of a file | `head <filename>` |
| `tail`  | Display the last 10 lines of a file | `tail <filename>` |
| `touch` | Create an empty file | `touch <filename>` |
| `mkdir` | Create a new directory | `mkdir <directory>` |
| `rm`    | Remove a file | `rm <filename>` |
| `cp`    | Copy a file | `cp <source> <destination>` |
| `mv`    | Move or rename a file | `mv <source> <destination>` |
| `echo`  | Print text to the console | `echo [text]` |
| `whoami`| Display the current username | `whoami` |
| `exit`  | Exit the shell | `exit` |

## Code Structure

- **`main.go`**: Contains the main entry point, the REPL loop, and implementations for all supported commands.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any enhancements or bug fixes.

## Inspiration

https://github.com/popovicu/ultimate-linux/