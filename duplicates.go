package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Asking the user to insert string
func AskForInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter any amount of numbers (separated by a space): ")
	text, _ := reader.ReadString('\n')
	return text
}

// Counting the duplicate strings
func CountDuplicates(str string) map[string]int {
	list := strings.Fields(str)
	duplicates := make(map[string]int)
	for _, word := range list {
		_, ok := duplicates[word]
		if ok {
			duplicates[word]++
		} else {
			duplicates[word] = 1
		}
	}
	return duplicates
}
func main() {
	count := AskForInput()
	fmt.Println("The characters duplicated are: ")

	for index, element := range CountDuplicates(count) {
		fmt.Println(index, "=>", element)
	}
}
