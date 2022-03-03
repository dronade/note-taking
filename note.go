package main

import (
	"bufio"
	"fmt"
	"os"
)

func createNote() {
	// be able to create note name
	// multiple line notes?
	file, err := os.Create("note3.txt")
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
}

func removeNote() {
	err := os.Remove("note3.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
}
