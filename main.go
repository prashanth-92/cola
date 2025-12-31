package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}

		cmd := parts[0]
		args := parts[1:]

		switch cmd {
		case "exit":
			return
		case "ls":
			handleLs(args)
		case "cat":
			handleCat(args)
		case "cd":
			handleCd(args)
		case "pwd":
			handlePwd()
		case "touch":
			handleTouch(args)
		case "mkdir":
			handleMkdir(args)
		case "rm":
			handleRm(args)
		case "cp":
			handleCp(args)
		case "mv":
			handleMv(args)
		case "echo":
			handleEcho(args)
		case "head":
			handleHead(args)
		case "tail":
			handleTail(args)
		case "whoami":
			handleWhoami()
		default:
			fmt.Printf("Unknown command: %s\n", cmd)
		}
	}
}

func handleLs(args []string) {
	dir := "."
	if len(args) > 0 {
		dir = args[0]
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading directory:", err)
		return
	}

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		if entry.IsDir() {
			fmt.Printf("%s/\n", entry.Name())
		} else {
			fmt.Printf("%s\t%d\n", entry.Name(), info.Size())
		}
	}
}

func handleCat(args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Usage: cat <filename>")
		return
	}

	for _, filename := range args {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", filename, err)
			continue
		}
		fmt.Print(string(data))
	}
}

func handleCd(args []string) {
	if len(args) == 0 {
		// In a real shell, this goes to HOME, but for now just usage or ignore
		fmt.Fprintln(os.Stderr, "Usage: cd <directory>")
		return
	}
	err := os.Chdir(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error changing directory:", err)
	}
}

func handlePwd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting working directory:", err)
		return
	}
	fmt.Println(dir)
}

func handleTouch(args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Usage: touch <filename>")
		return
	}
	for _, filename := range args {
		// Open for appending, which updates time if exists, or create if not.
		// Actually touch updates access/mod times. Simplest "touch": open and close.
		f, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error touching %s: %v\n", filename, err)
			continue
		}
		f.Close()
		// To be more accurate we should set current time, but this is "basic".
		now := time.Now()
		os.Chtimes(filename, now, now)
	}
}

func handleMkdir(args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Usage: mkdir <directory>")
		return
	}
	for _, dir := range args {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating directory %s: %v\n", dir, err)
		}
	}
}

func handleRm(args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Usage: rm <filename>")
		return
	}
	for _, filename := range args {
		err := os.Remove(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error removing %s: %v\n", filename, err)
		}
	}
}

func handleCp(args []string) {
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: cp <source> <dest>")
		return
	}
	src, dst := args[0], args[1]

	input, err := os.ReadFile(src)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading source:", err)
		return
	}

	err = os.WriteFile(dst, input, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error writing destination:", err)
	}
}

func handleMv(args []string) {
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: mv <source> <dest>")
		return
	}
	err := os.Rename(args[0], args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error moving file:", err)
	}
}

func handleEcho(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func handleHead(args []string) {
	printLines(args, 10, true)
}

func handleTail(args []string) {
	printLines(args, 10, false)
}

func printLines(args []string, n int, head bool) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Usage: head/tail <filename>")
		return
	}

	filename := args[0]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if head {
		if len(lines) < n {
			n = len(lines)
		}
		for i := 0; i < n; i++ {
			fmt.Println(lines[i])
		}
	} else {
		start := len(lines) - n
		if start < 0 {
			start = 0
		}
		for i := start; i < len(lines); i++ {
			fmt.Println(lines[i])
		}
	}
}

func handleWhoami() {
	user := os.Getenv("USER")
	if user == "" {
		fmt.Println("unknown")
	} else {
		fmt.Println(user)
	}
}
