package main

import (
	"bufio"
	"fmt"
	"os"
)

func createNote() {
	file, err := os.Create("note1.txt")
	if err != nil {
		fmt.Println(err)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter note: ")
	text, _ := reader.ReadString('\n')

	input, err := file.WriteString(text)
	if err != nil {
		fmt.Println(err)
	}
	print(input)

	file.Close()
	content, err := os.ReadFile("note1.txt")
	fmt.Print(content)
}

func main() {
	createNote()
}
