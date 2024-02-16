package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var fileName string = "log.txt"
var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
var date string = time.Now().Format("Monday, 02-01-2006 15:04:05")

func search_logs(f *os.File) {
	fmt.Print("Enter date: ")

	if !scanner.Scan() {
		fmt.Println("Error reading input:", scanner.Err())
		return
	}
	searchVal := scanner.Text()

	found := false
	var fetched string
	fScanner := bufio.NewScanner(f)
	for fScanner.Scan() {
		text := fScanner.Text()
		if found && text == " " || text == "------------------------------------------" {
			fetched += "\n"
			found = false
		}
		if found {
			fetched += text
		}
		if strings.Contains(text, searchVal) {
			fetched += text[len(text)-8:]
			fetched += "\n"
			found = true
		}
	}
	fmt.Print(fetched)
	fmt.Print("\n\n")
}

func add_log(f *os.File) {
	fmt.Printf("Your thoughts today for %s?\n", date)

	if !scanner.Scan() {
		fmt.Println("Error reading input:", scanner.Err())
		return
	}
	input := scanner.Text()
	fmt.Fprint(f, "------------------------------------------\n")
	fmt.Fprintf(f, "%s\n", date)

	lineLength := 40
	if len(input) <= lineLength {
		fmt.Fprintf(f, "-%s\n", input)
	} else {
		fmt.Fprint(f, "-")
		for i := 0; i < len(input); i += lineLength {
			end := i + lineLength
			if end > len(input) {
				end = len(input)
			}
			fmt.Fprintf(f, "%s\n", input[i:end])
		}
	}
	print("Have a nice day! :)")
}

func main() {
	fmt.Printf("Today is %s\n", date)
	fmt.Println("------------------------------------------")

	run := true
	for run {
		f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			panic(err)
		}

		fmt.Println("Log or Search?")
		fmt.Print("> ")

		if !scanner.Scan() {
			fmt.Println("Error reading input:", scanner.Err())
			return
		}
		input := scanner.Text()
		
		switch cmd := strings.ToLower(input); cmd {
		case "search":
			search_logs(f)
		case "log":
			add_log(f)
		case "exit":
			fmt.Println("Goodbye!")
			run = false
		case "help":
			fmt.Println("Use SEARCH, LOG, or EXIT")
		default:
			fmt.Println("I can only SEARCH or LOG :(")
		}
		defer f.Close()
	}
}
