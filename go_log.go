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
var date string = time.Now().Format("Monday, 02-01-2006")

func search_logs(f *os.File) {
	fmt.Print("Enter date: ")
	if scanner.Scan() {
		searchVal := scanner.Text()
		scanner = bufio.NewScanner(f)
		found := false
		for scanner.Scan() {
			text := scanner.Text()
			if found {
				fmt.Println(text)
				found = false
			}
			if strings.Contains(text, searchVal) {
				found = true
			}
		}
	}
}

func add_log(f *os.File) {
	fmt.Printf("Your thoughts today for %s?\n", date)
	if scanner.Scan() {
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
}

func main() {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Today is %s\n", date)
	fmt.Println("------------------------------------------")
	fmt.Println("Log or Search?")

	if scanner.Scan() {
		input := scanner.Text()
		if strings.ToLower(input) == "search" {
			search_logs(f)
		} else if strings.ToLower(input) == "log" {
			add_log(f)
		} else {
			fmt.Println("I can only SEARCH or LOG :(")
		}
	}
	defer f.Close()
}
