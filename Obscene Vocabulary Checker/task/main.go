package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	const exit = "exit"
	filename := InputFromUser()
	words := ReadFileAsWords(filename)
	for true {
		// write your code here

		sentence := InputFromUser()
		if sentence == exit {
			fmt.Println("Bye!")
			return
		}
		userWords := strings.Split(sentence, " ")
		for i, userWord := range userWords {
			if strings.Contains(words, strings.ToLower(userWord)) {
				fmt.Print(censorTabooWord(userWord))
			} else {
				fmt.Print(userWord)
			}
			if i == len(userWord)-1 {
				fmt.Println()
			} else {
				fmt.Print(" ")
			}
		}
	}
}

func censorTabooWord(taboo string) string {
	const censor = "*"
	return strings.Repeat(censor, len(taboo))
}

func ReadFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // this line closes the file before exiting the program

	scanner := bufio.NewScanner(file) // create a new Scanner for the file

	firstLine := true
	for scanner.Scan() {
		if firstLine {
			//fmt.Printf("The %s contents:\n", path)
			firstLine = false
		}
		fmt.Println(scanner.Text()) // the Text() function converts the scanned bytes to a string
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func ReadFileAsWords(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // this line closes the file before exiting the program

	scanner := bufio.NewScanner(file) // create a new Scanner for the file
	scanner.Split(bufio.ScanWords)

	firstLine := true
	words := ""
	for scanner.Scan() {
		if firstLine {
			//fmt.Printf("The %s contents:\n", path)
			firstLine = false
		}
		words = words + " " + strings.ToLower(scanner.Text())
	}
	return words
}

func InputFromUser() string {
	var input string
	fmt.Scan(&input)
	return input
}
