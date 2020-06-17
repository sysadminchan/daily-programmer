package main

import (
	"bufio"
	"fmt"
	"os"
)

func prompt() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter any amount of x's: ")
	text, _ := reader.ReadString('\n')
	fmt.Print("Enter any amount of y's: ")
	text2, _ := reader.ReadString('\n')
	return text, text2
}

func balanced(text, text2 string) bool {
	if len(text) == len(text2) {
		return true
	} else {
		return false
	}
}

func main() {
	text, text2 := prompt()
	fmt.Printf("You've entered the following strings: %s & %s", text, text2)
	fmt.Printf("Was it balanced? : %t\n", balanced(text, text2))
}
